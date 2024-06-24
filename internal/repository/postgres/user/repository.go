package user

import (
	"github.com/Frylock-dev/users/internal/model"
	"github.com/Frylock-dev/users/internal/repository"
	"github.com/Frylock-dev/users/internal/repository/postgres/user/converter"
	repoModel "github.com/Frylock-dev/users/internal/repository/postgres/user/model"
	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

type Repository struct {
	conn   *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewRepository(conn *pgxpool.Pool) repository.User {
	return &Repository{
		conn:   conn,
		getter: trmpgx.DefaultCtxGetter,
	}
}

func (repo *Repository) Save(ctx context.Context, uuid string, userInfo *model.UserInfo) error {
	txManager := manager.Must(trmpgx.NewDefaultFactory(repo.conn))

	if err := txManager.Do(ctx, func(ctx context.Context) error {
		id, err := repo.SaveUser(ctx, uuid)
		if err != nil {
			return err
		}

		err = repo.SaveUserInfo(ctx, id, userInfo)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) SaveUser(ctx context.Context, uuid string) (int, error) {
	var userID int

	conn := repo.getter.DefaultTrOrDB(ctx, repo.conn)

	query, args, err := sq.
		Insert("users").
		Columns("uuid").
		Values(uuid).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING \"users\".\"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	err = conn.
		QueryRow(ctx, query, args...).
		Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (repo *Repository) SaveUserInfo(ctx context.Context, userID int, userInfoService *model.UserInfo) error {
	conn := repo.getter.DefaultTrOrDB(ctx, repo.conn)
	userInfo := converter.ToUserInfoFromService(userInfoService)

	query, args, err := sq.
		Insert("users_info").
		Columns(
			"user_id",
			"phone",
			"email",
			"first_name",
			"second_name",
			"last_name",
			"passport_number",
			"passport_code",
			"passport_issue_date",
			"birthday",
		).Values(
		&userID,
		&userInfo.Phone,
		&userInfo.Email,
		&userInfo.FirstName,
		&userInfo.SecondName,
		&userInfo.LastName,
		&userInfo.PassportNumber,
		&userInfo.PassportCode,
		&userInfo.PassportIssueDate,
		&userInfo.Birthday,
	).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetByUUID(ctx context.Context, uuid string) (*model.User, error) {
	var user repoModel.User
	var userInfo repoModel.UserInfo

	conn := repo.getter.DefaultTrOrDB(ctx, repo.conn)

	query, args, err := sq.
		Select(
			"id",
			"uuid",
			"phone",
			"email",
			"first_name",
			"second_name",
			"last_name",
			"passport_number",
			"passport_code",
			"passport_issue_date",
			"birthday",
		).
		From("users").
		Join("users_info on users.id = users_info.user_id").
		Where(sq.Eq{"users.uuid": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.UUID,
		&userInfo.Phone,
		&userInfo.Email,
		&userInfo.FirstName,
		&userInfo.SecondName,
		&userInfo.LastName,
		&userInfo.PassportNumber,
		&userInfo.PassportCode,
		&userInfo.PassportIssueDate,
		&userInfo.Birthday,
	)
	if err != nil {
		return nil, err
	}

	user.Info = &userInfo

	return converter.ToUserFromRepo(&user), nil
}
