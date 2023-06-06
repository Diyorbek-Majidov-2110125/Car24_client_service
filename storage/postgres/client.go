package postgres

import (
	"CAR24/Car24_client_service/genproto/client_service"
	"CAR24/Car24_client_service/models"
	"CAR24/Car24_client_service/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (c *clientRepo) Create(ctx context.Context, req *client_service.CreateClientRequest) (resp *client_service.ClientPrimaryKey, err error) {
	id := uuid.New().String()

	query := `
		INSERT INTO client (id, first_name, last_name, address, phone_number, driving_license_number, passport_number, photo, driving_number_given_place, driving_number_given_date, driving_number_given_expire, date_of_birth,  propiska,  passport_pinfl,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, NOW())
		`

	_, err = c.db.Exec(ctx, query, id, req.FirstName, req.LastName, req.Address, req.PhoneNumber, req.DrivingLicenseNumber, req.PassportNumber, req.Photo, req.DrivingNumberGivenPlace, req.DrivingNumberGivenDate, req.DrivingNumberGivenExpire, req.DateOfBirth, req.Propiska, req.PassportPinfl,)
	if err != nil {
		return nil, err
	}

	return &client_service.ClientPrimaryKey{Id: id}, nil
}

func (c *clientRepo) GetById(ctx context.Context, req *client_service.ClientPrimaryKey) (client *client_service.Client, err error) {
	fmt.Println("This is here")
	query := `
		SELECT id, first_name, last_name, address, phone_number, driving_license_number, passport_number, photo, created_at, updated_at, driving_number_given_place, driving_number_given_date, driving_number_given_expire, date_of_birth, is_blocked, propiska, passport_pinfl
		FROM client
		WHERE id = $1
	`
	fmt.Println("hi")

	var (
		id                       sql.NullString
		firstName                sql.NullString
		lastName                 sql.NullString
		address                  sql.NullString
		phoneNumber              sql.NullString
		drivingLicenseNumber     sql.NullString
		passportNumber           sql.NullString
		photo                    sql.NullString
		createdAt                sql.NullString
		updatedAt                sql.NullString
		drivingNumberGivenPlace  sql.NullString
		drivingNumberGivenDate   sql.NullString
		drivingNumberGivenExpire sql.NullString
		dateOfBirth              sql.NullString
		isBlocked                sql.NullBool
		propiska                 sql.NullString
		passportPinfl            sql.NullString
	)
	row := c.db.QueryRow(ctx, query, req.Id)
fmt.Println("before err")
	err = row.Scan(&id, &firstName, &lastName, &address, &phoneNumber, &drivingLicenseNumber, &passportNumber, &photo, &createdAt, &updatedAt, &drivingNumberGivenPlace, &drivingNumberGivenDate, &drivingNumberGivenExpire, &dateOfBirth, &isBlocked, &propiska, &passportPinfl)
	if err != nil {
		fmt.Println("good")
		return nil, err
	}

	client = &client_service.Client{
		Id:                       id.String,
		FirstName:                firstName.String,
		LastName:                 lastName.String,
		Address:                  address.String,
		PhoneNumber:              phoneNumber.String,
		DrivingLicenseNumber:     drivingLicenseNumber.String,
		PassportNumber:           passportNumber.String,
		Photo:                    photo.String,
		CreatedAt:                createdAt.String,
		UpdatedAt:                updatedAt.String,
		DrivingNumberGivenPlace:  drivingNumberGivenPlace.String,
		DrivingNumberGivenDate:   drivingNumberGivenDate.String,
		DrivingNumberGivenExpire: drivingNumberGivenExpire.String,
		DateOfBirth:              dateOfBirth.String,
		IsBlocked:                isBlocked.Bool,
		Propiska:                 propiska.String,
		PassportPinfl:            passportPinfl.String,
	}

	return
}

func (c *clientRepo) Update(ctx context.Context, req *client_service.UpdateClient) (rowsaffected int64, err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE client
		SET id = :id,
			first_name = :first_name,
			last_name = :last_name,
			address = :address,
			phone_number = :phone_number,
			driving_license_number = :driving_license_number,
			passport_number = :passport_number,
			photo = :photo,
			updated_at = NOW(),
			driving_number_given_place = :driving_number_given_place,
			driving_number_given_date = :driving_number_given_date,
			driving_number_given_expire = :driving_number_given_expire,
			date_of_birth = :date_of_birth,
			is_blocked = :is_blocked,
			propiska = :propiska,
			passport_pinfl = :passport_pinfl
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":                          req.GetId(),
		"first_name":                  req.GetFirstName(),
		"last_name":                   req.GetLastName(),
		"address":                     req.GetAddress(),
		"phone_number":                req.GetPhoneNumber(),
		"driving_license_number":      req.GetDrivingLicenseNumber(),
		"passport_number":             req.GetPassportNumber(),
		"photo":                       req.GetPhoto(),
		"driving_number_given_place":  req.GetDrivingNumberGivenPlace(),
		"driving_number_given_date":   req.GetDrivingNumberGivenDate(),
		"driving_number_given_expire": req.GetDrivingNumberGivenExpire(),
		"date_of_birth":               req.GetDateOfBirth(),
		"is_blocked":                  req.GetIsBlocked(),
		"propiska":                    req.GetPropiska(),
		"passport_pinfl":              req.GetPassportPinfl(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (c *clientRepo) GetList(ctx context.Context, req *client_service.GetListClientRequest) (resp *client_service.GetListClientResponse, err error) {
	resp = &client_service.GetListClientResponse{}

	var (
		query  string
		limit  = ""
		offset = " offset 0 "
		params = make(map[string]interface{})
		filter = " where true	"
		sort   = " order by created_at desc	 "
	)

	query = `
		select id, first_name, last_name, address, phone_number, driving_license_number, passport_number, photo, created_at, updated_at, driving_number_given_place, driving_number_given_date, driving_number_given_expire, date_of_birth, is_blocked, propiska, passport_pinfl
		from client
		`
	if len(req.GetSearch()) > 0 {
		filter += " and (first_name ilike :search or last_name ilike :search or phone_number ilike :search or passport_number ilike :search or driving_license_number ilike :search)"
		params["search"] = "%" + req.GetSearch() + "%"
	}

	if req.GetLimit() > 0 {
		limit = " LIMIT :	limit"
		params["limit"] = req.Limit
	}

	if req.GetOffset() > 0 {
		offset = " OFFSET :	offset"
		params["offset"] = req.Offset
	}

	query += filter + sort + limit + offset

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id                       sql.NullString
			firstName                sql.NullString
			lastName                 sql.NullString
			address                  sql.NullString
			phoneNumber              sql.NullString
			drivingLicenseNumber     sql.NullString
			passportNumber           sql.NullString
			photo                    sql.NullString
			createdAt                sql.NullString
			updatedAt                sql.NullString
			drivingNumberGivenPlace  sql.NullString
			drivingNumberGivenDate   sql.NullString
			drivingNumberGivenExpire sql.NullString
			dateOfBirth              sql.NullString
			isBlocked                sql.NullBool
			propiska                 sql.NullString
			passportPinfl            sql.NullString
		)

		err := rows.Scan(
			&id,
			&firstName,
			&lastName,
			&address,
			&phoneNumber,
			&drivingLicenseNumber,
			&passportNumber,
			&photo,
			&createdAt,
			&updatedAt,
			&drivingNumberGivenPlace,
			&drivingNumberGivenDate,
			&drivingNumberGivenExpire,
			&dateOfBirth,
			&isBlocked,
			&propiska,
			&passportPinfl,
		)

		if err != nil {
			return nil, err
		}

		resp.Clients = append(resp.Clients, &client_service.Client{
			Id:                       id.String,
			FirstName:                firstName.String,
			LastName:                 lastName.String,
			Address:                  address.String,
			PhoneNumber:              phoneNumber.String,
			DrivingLicenseNumber:     drivingLicenseNumber.String,
			PassportNumber:           passportNumber.String,
			Photo:                    photo.String,
			CreatedAt:                createdAt.String,
			UpdatedAt:                updatedAt.String,
			DrivingNumberGivenPlace:  drivingNumberGivenPlace.String,
			DrivingNumberGivenDate:   drivingNumberGivenDate.String,
			DrivingNumberGivenExpire: drivingNumberGivenExpire.String,
			DateOfBirth:              dateOfBirth.String,
			IsBlocked:                isBlocked.Bool,
			Propiska:                 propiska.String,
			PassportPinfl:            passportPinfl.String,
		})
	}
	return
}

func (c *clientRepo) UpdatePatch(ctx context.Context, req models.UpdatePatchRequest) (rowsaffected int64, err error) {
	var (
		set   = " Set "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key:= range req.Fields {
		set += fmt.Sprintf(" %s = :%s", key, key)
		if ind != len(req.Fields)-1 {
			set += ","
		}
		ind++
	}

	query = `
		update client
		` + set + ` , updated_at = now()
		where id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (c *clientRepo) Delete(ctx context.Context, req *client_service.ClientPrimaryKey) (err error) {
	query := `
		delete from client
		where id = $1
	`

	_ , err = c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
