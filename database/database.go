package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JeF11037/simple_web_goapp/models"

	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("\nError: %s", err)
	}
}

func Connection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("connection"))
	CheckError(err)
	return db
}

func Ping() bool {
	db := Connection()
	defer db.Close()
	err := db.Ping()
	CheckError(err)
	return err == nil
}

func CheckUserTable() bool {
	db := Connection()
	defer db.Close()
	_, err := db.Exec(`
		SELECT * FROM public.user;
	`)
	CheckError(err)
	return err == nil
}

func CreateUserTable() bool {
	db := Connection()
	defer db.Close()
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS public.user (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		birth_date DATE NOT NULL,
		gender VARCHAR(20) NOT NULL,
		email VARCHAR(100) NOT NULL,
		address VARCHAR(200)
	)
	`)
	CheckError(err)
	return err == nil
}

func FillSeedData() bool {
	db := Connection()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Pamela', 'D. Heideman', '1984-10-31', 'Female', 'PamelaDHeideman@rhyta.com', '3520 Bingamon Branch Road');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Melissa', 'M. Franklin', '2000-07-29', 'Female', 'MelissaMFranklin@jourrapide.com ', '569 Burning Memory Lane');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Merle', 'R. Britt', '1979-04-11', 'Female', 'MerleRBritt@jourrapide.com ', '1231 Cherry Tree Drive');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Clayton', 'C. Bowman', '1978-09-01', 'Male', 'ClaytonCBowman@dayrep.com', '139 Heavner Avenue');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('George', 'M. Foy', '2002-12-25', 'Male', 'GeorgeMFoy@dayrep.com ', '4964 Charmaine Lane');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Robert', 'L. Tucci', '1983-09-24', 'Male', 'RobertLTucci@armyspy.com ', '2512 North Street');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('John', 'W. Williams', '1993-02-09', 'Male', 'JohnWWilliams@armyspy.com', '1934 Jenna Lane');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('David', 'R. Ruby', '1971-04-16', 'Male', 'DavidRRuby@jourrapide.com ', '1130 Frank Avenue');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Helga', 'I. Taylor', '1995-08-20', 'Female', 'HelgaITaylor@teleworm.us ', '2279 Layman Avenue');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Misty', 'S. Cohn', '1992-05-26', 'Female', 'MistySCohn@teleworm.us ', '1872 Mulberry Lane');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Lisa', 'J. Montgomery', '1973-06-14', 'Female', 'LisaJMontgomery@jourrapide.com ', '2441 Eagle Drive');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Karl', 'J. Fleming', '1972-10-04', 'Male', 'KarlJFleming@teleworm.us ', '60 Lonely Oak Drive');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Marie', 'D. Kidd', '1996-06-04', 'Female', '4952 Duffy Street', 'MarieDKidd@armyspy.com ');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Richard', 'M. Moffat', '1983-07-14', 'Male', 'RichardMMoffat@rhyta.com ', '3432 Kessla Way');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Hope', 'J. Robinson', '1981-01-29', 'Female', 'HopeJRobinson@rhyta.com ', '1441 Lowndes Hill Park Road');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Erica', 'O. Dugger', '2003-07-06', 'Female', 'EricaODugger@jourrapide.com ', '4632 Hickman Street');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Bradley', 'A. Sanders', '1991-12-29', 'Male', 'BradleyASanders@teleworm.us ', '701 Pooz Street');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Dorothy', 'R. Hollins', '1995-05-28', 'Female', 'DorothyRHollins@rhyta.com ', '2728 Emerson Road');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Dana', 'A. Apgar', '1984-01-25', 'Male', 'DanaAApgar@rhyta.com ', '3116 Flanigan Oaks Drive');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Gertrude', 'A. Rose', '1960-06-28', 'Female', 'GertrudeARose@teleworm.us ', '1655 New York Avenue');
		INSERT INTO public.user(first_name, last_name, birth_date, gender, email, address)
		VALUES('Allison', 'H. Davis', '1966-07-17', 'Female', 'AllisonHDavis@teleworm.us ', '4208 Rardin Drive');
	`)
	CheckError(err)
	return err == nil
}

func GetUsers() []models.User {
	var users []models.User
	var user models.User
	db := Connection()
	defer db.Close()
	rows, err := db.Query(`
	SELECT
	id, first_name, last_name,
	TO_CHAR(birth_date :: DATE, 'yyyy-mm-dd'),
	gender, email, address
	FROM public.user
	`)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.BirthDate,
			&user.Gender,
			&user.Email,
			&user.Address)
		CheckError(err)
		users = append(users, user)
	}
	err = rows.Err()
	CheckError(err)
	return users
}

func AddUser(user models.User) bool {
	db := Connection()
	defer db.Close()
	_, err := db.Query(`
	INSERT INTO public.user (
		first_name,
		last_name,
		birth_date,
		gender,
		email,
		address)
		VALUES ($1, $2, $3, $4, $5, $6)
	`,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Gender,
		user.Email,
		user.Address)
	CheckError(err)
	return err == nil
}

func UpdateUser(user models.User) bool {
	db := Connection()
	defer db.Close()
	_, err := db.Exec(`
	UPDATE public.user
	SET
	first_name = $1,
	last_name = $2,
	birth_date = $3,
	gender = $4,
	email = $5,
	address = $6
	WHERE id = $7
	`,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Gender,
		user.Email,
		user.Address,
		user.ID)
	CheckError(err)
	return err == nil
}

func DeleteUser(user models.User) bool {
	db := Connection()
	defer db.Close()
	_, err := db.Exec(`
	DELETE FROM public.user
	WHERE id = $1
	`,
		user.ID)
	CheckError(err)
	return err == nil
}
