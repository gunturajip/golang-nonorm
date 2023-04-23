package main

import (
	"encoding/json"
	"golang-nonorm/database"
	"golang-nonorm/helpers"
	"golang-nonorm/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	database.StartDB()
	defer database.CloseDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/mahasiswa", GetAllMahasiswa).Methods("GET")
	router.HandleFunc("/api/v1/mahasiswa", PostMahasiswa).Methods("POST")
	router.HandleFunc("/api/v1/mahasiswa/{id}", GetMahasiswa).Methods("GET")
	router.HandleFunc("/api/v1/mahasiswa/{id}", UpdateMahasiswa).Methods("PUT")
	router.HandleFunc("/api/v1/mahasiswa/{id}", DeleteMahasiswa).Methods("DELETE")

	serverport := ":8080"
	err := http.ListenAndServe(serverport, router)
	if err != nil {
		log.Fatal("error in server", err)
	}
}

// var mahasiswa = []models.Mahasiswa{
// 	{Id: 1, Nama: "guntur", Usia: 20, Gender: false, TanggalRegistrasi: time.Now()},
// 	{Id: 2, Nama: "shofy", Usia: 21, Gender: true, TanggalRegistrasi: time.Now()},
// 	{Id: 3, Nama: "aji", Usia: 22, Gender: false, TanggalRegistrasi: time.Now()},
// 	{Id: 4, Nama: "yorinda", Usia: 23, Gender: true, TanggalRegistrasi: time.Now()},
// 	{Id: 5, Nama: "pratama", Usia: 24, Gender: false, TanggalRegistrasi: time.Now()},
// }

func GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	// set response header into application/json
	w.Header().Set("Content-Type", "application/json")

	// open up database, define sql query, and prepare to catch internal server error
	db := database.GetDB()
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := helpers.Response{
				Status: 500,
				Error:  "internal server error",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}()

	// execute sql query for getting all mahasiswa data including nama_jurusan and nama_hobi
	mjhs := []helpers.MJH{}
	sqlStat := `SELECT * FROM mahasiswa`
	rows, err := db.Query(sqlStat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid parameter while reading all mahasiswa data",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// store result of sql query to slice of mahasiswa data including nama_jurusan and nama_hobi
	var iter int
	for rows.Next() {
		var mjh = helpers.MJH{}
		var tanggalRegistrasi helpers.TimeScanner
		err = rows.Scan(&mjh.Mahasiswa.ID, &mjh.Mahasiswa.Nama, &mjh.Mahasiswa.Usia, &mjh.Mahasiswa.Gender, &tanggalRegistrasi)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid argument while reading a mahasiswa data",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		mjh.Mahasiswa.TanggalRegistrasi = tanggalRegistrasi.Time
		sqlStat := `
			SELECT nama_jurusan, nama_hobi FROM mahasiswa_jurusan
			JOIN mahasiswa_hobi ON mahasiswa_jurusan.id_mahasiswa = mahasiswa_hobi.id_mahasiswa
			JOIN jurusan ON mahasiswa_jurusan.id_jurusan = jurusan.id
			JOIN hobi ON mahasiswa_hobi.id_hobi = hobi.id
			WHERE mahasiswa_jurusan.id_mahasiswa = ?
		`
		jur_hob, err := db.Query(sqlStat, mjh.Mahasiswa.ID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid parameter while reading a mahasiswa_jurusan, mahasiswa_hobi, jurusan, and hobi data",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		for jur_hob.Next() {
			err = jur_hob.Scan(&mjh.NamaJurusan, &mjh.NamaHobi)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response := helpers.Response{
					Status: 400,
					Error:  "invalid argument while reading a mahasiswa_jurusan, mahasiswa_hobi, jurusan, and hobi data",
				}
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		mjhs = append(mjhs, mjh)
		iter++
	}

	// check if the query returns no rows
	if iter == 0 {
		w.WriteHeader(http.StatusNotFound)
		response := helpers.Response{
			Status: 404,
			Error:  "mahasiswa data not found",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// return succes response
	w.WriteHeader(http.StatusOK)
	response := helpers.Response{
		Status:  200,
		Success: "successfully get all mahasiswa data including nama_jurusan and nama_hobi",
		Data:    mjhs,
	}
	json.NewEncoder(w).Encode(response)
}

// func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
// studentId := r.URL.Query().Get("id")
// repository.GetStudents(studentId)
// }

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	// set response header into application/json
	w.Header().Set("Content-Type", "application/json")

	// open up database, define sql query, and prepare to catch internal server error
	db := database.GetDB()
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := helpers.Response{
				Status: 500,
				Error:  "internal server error",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}()

	// execute sql query for getting a mahasiswa data
	idMhs, _ := strconv.Atoi(mux.Vars(r)["id"])
	var mjh = helpers.MJH{}
	sqlStat := `SELECT * FROM mahasiswa WHERE id = ?`
	mhs, err := db.Query(sqlStat, idMhs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid argument while reading a mahasiswa data",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// continue to include nama_jurusan and nama_hobi of the mahasiswa
	var iter int
	for mhs.Next() {
		var tanggalRegistrasi helpers.TimeScanner
		err = mhs.Scan(&mjh.Mahasiswa.ID, &mjh.Mahasiswa.Nama, &mjh.Mahasiswa.Usia, &mjh.Mahasiswa.Gender, &tanggalRegistrasi)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid argument while reading a mahasiswa data",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		mjh.Mahasiswa.TanggalRegistrasi = tanggalRegistrasi.Time
		sqlStat := `
			SELECT nama_jurusan, nama_hobi FROM mahasiswa_jurusan
			JOIN mahasiswa_hobi ON mahasiswa_jurusan.id_mahasiswa = mahasiswa_hobi.id_mahasiswa
			JOIN jurusan ON mahasiswa_jurusan.id_jurusan = jurusan.id
			JOIN hobi ON mahasiswa_hobi.id_hobi = hobi.id
			WHERE mahasiswa_jurusan.id_mahasiswa = ?
		`
		jur_hob, err := db.Query(sqlStat, mjh.Mahasiswa.ID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid parameter while reading a mahasiswa_jurusan, mahasiswa_hobi, jurusan, and hobi data",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		for jur_hob.Next() {
			err = jur_hob.Scan(&mjh.NamaJurusan, &mjh.NamaHobi)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response := helpers.Response{
					Status: 400,
					Error:  "invalid argument while reading a mahasiswa_jurusan, mahasiswa_hobi, jurusan, and hobi data",
				}
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		iter++
	}

	// check if the query returns no rows
	if iter == 0 {
		w.WriteHeader(http.StatusNotFound)
		response := helpers.Response{
			Status: 404,
			Error:  "mahasiswa data not found",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// return succes response
	w.WriteHeader(http.StatusOK)
	response := helpers.Response{
		Status:  200,
		Success: "successfully get a mahasiswa data including nama_jurusan and nama_hobi",
		Data:    mjh,
	}
	json.NewEncoder(w).Encode(response)
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	// set response header into application/json
	w.Header().Set("Content-Type", "application/json")

	// open up database, define sql query, and prepare to catch internal server error
	db := database.GetDB()
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := helpers.Response{
				Status: 500,
				Error:  "internal server error",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}()

	// validate all input from form data
	nama := r.FormValue("nama")
	usia := r.FormValue("usia")
	gender := r.FormValue("gender")
	tanggalRegistrasi := time.Now()
	nama_jurusan := r.FormValue("nama_jurusan")
	nama_hobi := r.FormValue("nama_hobi")
	if len(strings.TrimSpace(nama)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid nama in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(strings.TrimSpace(nama_jurusan)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid nama_jurusan in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(strings.TrimSpace(nama_hobi)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid nama_hobi in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	usiaConv, err := strconv.Atoi(usia)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid usia in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	genderConv, err := strconv.ParseBool(gender)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid gender in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// execute sql query for inserting a new mahasiswa data
	var mjhmjmh = helpers.MJHMJMH{}
	sqlStat := `
		INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi)
		VALUES (?, ?, ?, ?)
	`
	tmp, err := db.Exec(sqlStat, nama, usiaConv, genderConv, tanggalRegistrasi)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid argument while inserting mahasiswa data into database",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	idMhs, _ := tmp.LastInsertId()
	mhs := models.Mahasiswa{
		ID:                int(idMhs),
		Nama:              nama,
		Usia:              usiaConv,
		Gender:            genderConv,
		TanggalRegistrasi: tanggalRegistrasi,
	}
	mjhmjmh.Mahasiswa = mhs

	// execute sql query for checking the existence of nama_jurusan value in jurusan and insert it if not exist
	jur := models.Jurusan{}
	var idJur int64
	rows, _ := db.Query(`SELECT * FROM jurusan`)
	for rows.Next() {
		var j = models.Jurusan{}
		_ = rows.Scan(&j.ID, &j.NamaJurusan)
		if j.NamaJurusan == nama_jurusan {
			idJur = int64(j.ID)
			jur.ID = j.ID
			jur.NamaJurusan = nama_jurusan
		}
	}
	if jur.ID == 0 {
		sqlStat = `
			INSERT INTO jurusan (nama_jurusan)
			VALUES (?)
		`
		tmp, err = db.Exec(sqlStat, nama_jurusan)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid argument while inserting jurusan data into database",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		idJur, _ = tmp.LastInsertId()
		jur.ID = int(idJur)
		jur.NamaJurusan = nama_jurusan
	}
	mjhmjmh.Jurusan = jur

	// execute sql query for checking the existence of nama_hobi value in database and insert it if not exist
	hob := models.Hobi{}
	var idHob int64
	rows, _ = db.Query(`SELECT * FROM hobi`)
	for rows.Next() {
		var h = models.Hobi{}
		_ = rows.Scan(&h.ID, &h.NamaHobi)
		if h.NamaHobi == nama_hobi {
			idHob = int64(h.ID)
			hob.ID = h.ID
			hob.NamaHobi = nama_hobi
		}
	}
	if hob.ID == 0 {
		sqlStat = `
			INSERT INTO hobi (nama_hobi)
			VALUES (?)
		`
		tmp, err = db.Exec(sqlStat, nama_hobi)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid argument while inserting hobi data into database",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		idHob, _ = tmp.LastInsertId()
		hob.ID = int(idHob)
		hob.NamaHobi = nama_hobi
	}
	mjhmjmh.Hobi = hob

	// execute sql query for inserting a new mahasiswa_jurusan data
	sqlStat = `
		INSERT INTO mahasiswa_jurusan (id_mahasiswa, id_jurusan)
		VALUES (?, ?)
	`
	_, err = db.Exec(sqlStat, idMhs, idJur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid argument while inserting mahasiswa_jurusan data into database",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	mhs_jur := models.MahasiswaJurusan{
		IDMahasiswa: int(idMhs),
		IDJurusan:   int(idJur),
	}
	mjhmjmh.MahasiswaJurusan = mhs_jur

	// execute sql query for inserting a new mahasiswa_hobi data
	sqlStat = `
		INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi)
		VALUES (?, ?)
	`
	_, err = db.Exec(sqlStat, idMhs, idHob)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status:  400,
			Error:   "invalid argument while inserting mahasiswa_hobi data into database",
			Success: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	mhs_hob := models.MahasiswaHobi{
		IDMahasiswa: int(idMhs),
		IDHobi:      int(idHob),
	}
	mjhmjmh.MahasiswaHobi = mhs_hob

	// return succes response
	w.WriteHeader(http.StatusCreated)
	response := helpers.Response{
		Status:  201,
		Success: "successfully created data mahasiswa (include jurusan and hobi)",
		Data:    mjhmjmh,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// open up database, define sql query, and prepare to catch internal server error
	db := database.GetDB()
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := helpers.Response{
				Status: 500,
				Error:  "internal server error",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}()

	// validate all input from form data
	nama := r.FormValue("nama")
	usia := r.FormValue("usia")
	gender := r.FormValue("gender")
	nama_jurusan := r.FormValue("nama_jurusan")
	nama_hobi := r.FormValue("nama_hobi")
	if len(strings.TrimSpace(nama)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid nama in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(strings.TrimSpace(nama_jurusan)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid nama_jurusan in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(strings.TrimSpace(nama_hobi)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid nama_hobi in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	usiaConv, err := strconv.Atoi(usia)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid usia in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	genderConv, err := strconv.ParseBool(gender)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid gender in the input",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var mjhmjmh = helpers.MJHMJMH{}
	idMhs, _ := strconv.Atoi(mux.Vars(r)["id"])
	sqlStat := `
		UPDATE mahasiswa SET nama = ?, usia = ?, gender = ?
		WHERE id = ?
	`
	_, err = db.Exec(sqlStat, nama, usiaConv, genderConv, idMhs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid argument while updating mahasiswa data in database",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	mjhmjmh.Mahasiswa = models.Mahasiswa{
		ID:     idMhs,
		Nama:   nama,
		Usia:   usiaConv,
		Gender: genderConv,
	}

	// execute sql query for checking the existence of nama_jurusan value in jurusan and insert it if not exist
	sqlStat = `
		SELECT id, nama_jurusan, id_mahasiswa, id_jurusan FROM mahasiswa_jurusan
		JOIN jurusan ON mahasiswa_jurusan.id_jurusan = jurusan.id
		WHERE mahasiswa_jurusan.id_mahasiswa = ?
	`
	rows, _ := db.Query(sqlStat, idMhs)
	var iter int
	for rows.Next() {
		_ = rows.Scan(&mjhmjmh.Jurusan.ID, &mjhmjmh.Jurusan.NamaJurusan, &mjhmjmh.MahasiswaJurusan.IDMahasiswa, &mjhmjmh.MahasiswaJurusan.IDJurusan)
		iter++
	}
	if iter == 0 {
		sqlStat = `
			INSERT INTO jurusan (nama_jurusan)
			VALUES (?)
		`
		tmp, err := db.Exec(sqlStat, nama_jurusan)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid argument while inserting jurusan data into database",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		idJur, _ := tmp.LastInsertId()
		mjhmjmh.Jurusan = models.Jurusan{
			ID:          int(idJur),
			NamaJurusan: nama_jurusan,
		}

		// execute sql query for inserting a new mahasiswa_jurusan data
		sqlStat := `
			UPDATE mahasiswa_jurusan SET id_jurusan = ?
			WHERE id_mahasiswa = ?
		`
		_, err = db.Exec(sqlStat, idJur, idMhs)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status:  400,
				Error:   "invalid argument while updating mahasiswa_jurusan data in database",
				Success: err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		mjhmjmh.MahasiswaJurusan = models.MahasiswaJurusan{
			IDMahasiswa: idMhs,
			IDJurusan:   int(idJur),
		}
	}

	// execute sql query for checking the existence of nama_hobi value in hobi and insert it if not exist
	sqlStat = `
		SELECT id, nama_hobi, id_mahasiswa, id_hobi FROM mahasiswa_hobi
		JOIN hobi ON mahasiswa_hobi.id_hobi = hobi.id
		WHERE mahasiswa_hobi.id_mahasiswa = ?
	`
	rows, _ = db.Query(sqlStat, idMhs)
	iter = 0
	for rows.Next() {
		_ = rows.Scan(&mjhmjmh.Hobi.ID, &mjhmjmh.Hobi.NamaHobi, &mjhmjmh.MahasiswaHobi.IDMahasiswa, &mjhmjmh.MahasiswaHobi.IDHobi)
		iter++
	}
	if iter == 0 {
		sqlStat = `
			INSERT INTO hobi (nama_hobi)
			VALUES (?)
		`
		tmp, err := db.Exec(sqlStat, nama_hobi)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status: 400,
				Error:  "invalid argument while inserting hobi data into database",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		idHob, _ := tmp.LastInsertId()

		// execute sql query for inserting a new mahasiswa_hobi data
		sqlStat := `
			UPDATE mahasiswa_hobi SET id_hobi = ?
			WHERE id_mahasiswa = ?
		`
		_, err = db.Exec(sqlStat, idHob, idMhs)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := helpers.Response{
				Status:  400,
				Error:   "invalid argument while updating mahasiswa_hobi data in database",
				Success: err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// return succes response
	w.WriteHeader(http.StatusOK)
	response := helpers.Response{
		Status:  200,
		Success: "successfully updated data mahasiswa (include jurusan and hobi)",
		Data:    mjhmjmh,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// open up database, define sql query, and prepare to catch internal server error
	db := database.GetDB()
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := helpers.Response{
				Status: 500,
				Error:  "internal server error",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}()

	idMhs := mux.Vars(r)["id"]
	sqlStat := `SELECT * FROM mahasiswa WHERE id = ?`
	rows, _ := db.Query(sqlStat, idMhs)
	if !rows.Next() {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 404,
			Error:  "mahasiswa data not found",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	sqlStat = `
		DELETE FROM mahasiswa_jurusan
		WHERE id_mahasiswa = ?
	`
	_, err := db.Exec(sqlStat, idMhs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid argument while deleting mahasiswa_jurusan data in database",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	sqlStat = `
		DELETE FROM mahasiswa_hobi
		WHERE id_mahasiswa = ?
	`
	_, err = db.Exec(sqlStat, idMhs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status: 400,
			Error:  "invalid argument while deleting mahasiswa_hobi data in database",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	sqlStat = `
		DELETE FROM mahasiswa
		WHERE id = ?
	`
	_, err = db.Exec(sqlStat, idMhs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := helpers.Response{
			Status:  400,
			Error:   "invalid argument while deleting mahasiswa data in database",
			Success: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// return succes response
	w.WriteHeader(http.StatusOK)
	response := helpers.Response{
		Status:  200,
		Success: "successfully deleted data mahasiswa (including mahasiswa_jurusan, mahasiswa_hobi)",
	}
	json.NewEncoder(w).Encode(response)
}
