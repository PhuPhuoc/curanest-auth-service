package accountrepository

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/jmoiron/sqlx"
)

func (r *accountRepo) GetAccountWithFilter(ctx context.Context, filter *accountqueries.FilterAccountQuery) ([]accountdomain.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	whereClause, args := prepareWhereClause(&filter.Filter)

	errchan := make(chan error, 2)
	countchan := make(chan int, 1)
	datachan := make(chan []accountdomain.Account, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	go r.getCount(ctx, whereClause, args, errchan, countchan, &wg)
	go r.getData(ctx, filter.Paging, whereClause, args, errchan, datachan, &wg)

	var once sync.Once // Make sure to close the channel only once.
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errchan <- fmt.Errorf("panic in goroutine: %v", r)
			}
		}()

		wg.Wait()
		once.Do(func() {
			close(errchan)
			close(countchan)
			close(datachan)
		})
	}()

	var totalRecord int
	var accounts []accountdomain.Account

	receivedCount := 0
	expectedCount := 2

	for {
		select {
		case err, ok := <-errchan:
			if ok {
				return nil, err
			}
		case count, ok := <-countchan:
			if ok {
				totalRecord = count
				receivedCount++
			}
		case data, ok := <-datachan:
			if ok {
				accounts = data
				receivedCount++
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("operation timed out: %w", ctx.Err())
		}

		if receivedCount == expectedCount {
			break
		}
	}

	totalPages := int(math.Ceil(float64(totalRecord) / float64(filter.Paging.Size)))
	filter.Paging.Total = totalPages
	return accounts, nil
}

func (r *accountRepo) getCount(
	ctx context.Context,
	whereClause string,
	args map[string]interface{},
	errchan chan<- error,
	countchan chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	query := "select count(*) from " + TABLE + whereClause

	queryWithArgs, argList, err := sqlx.Named(query, args)
	if err != nil {
		errchan <- err
		return
	}
	queryWithArgs = r.db.Rebind(queryWithArgs)

	var total int
	if err := r.db.GetContext(ctx, &total, queryWithArgs, argList...); err != nil {
		errchan <- err
		return
	}

	countchan <- total
}

func (r *accountRepo) getData(
	ctx context.Context,
	paging common.Paging,
	whereClause string,
	args map[string]interface{},
	errchan chan<- error,
	datachan chan<- []accountdomain.Account,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	queryArgs := make(map[string]interface{})
	for k, v := range args {
		queryArgs[k] = v
	}
	queryArgs["limit"] = paging.Size
	queryArgs["offset"] = (paging.Page - 1) * paging.Size
	limit := " limit :limit offset :offset"

	order := " order by created_at desc"
	fields := strings.Join(GET_FIELD, ", ")
	query := "select " + fields + " from " + TABLE + whereClause + order + limit

	queryWithArgs, argList, err := sqlx.Named(query, queryArgs)
	if err != nil {
		errchan <- err
		return
	}
	queryWithArgs = r.db.Rebind(queryWithArgs)

	var listdto []AccountDTO
	if err := r.db.SelectContext(ctx, &listdto, queryWithArgs, argList...); err != nil {
		errchan <- err
		return
	}

	var entities []accountdomain.Account
	for _, dto := range listdto {
		entity, _ := dto.ToEntity()
		entities = append(entities, *entity)
	}

	datachan <- entities
}

func prepareWhereClause(param *accountqueries.FieldFilterAccount) (string, map[string]interface{}) {
	conditions := []string{}
	args := map[string]interface{}{}

	if param.RoleId != "" {
		conditions = append(conditions, "role_id = :role_id")
		args["role_id"] = param.RoleId
	}

	if param.FullName != "" {
		conditions = append(conditions, "full_name like :full_name")
		args["full_name"] = "%" + param.FullName + "%"
	}

	if param.Email != "" {
		conditions = append(conditions, "email like :email")
		args["email"] = "%" + param.Email + "%"
	}

	if param.PhoneNumber != "" {
		conditions = append(conditions, "phone_number like :phone_number")
		args["phone_number"] = "%" + param.PhoneNumber + "%"
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " where " + strings.Join(conditions, " and ")
	}

	return whereClause, args
}
