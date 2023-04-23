export $(grep -v '^#' .env | xargs)
# RESIM (Mahasiswa Simple REST API)

RESIM is an Instagram Clone for CRUD-ing mahasiswa data, as well as their jurusan and hobi on separate tables.
To-do:
- Need update for repository and service
- Change id to uuid in any models for data security
- Optimize all code for shorter code and faster calculation
- Make swagger documentation for easier readibility

## Project Setup

To run the project locally, follow these steps:

1. Clone the repository.
2. Install the required dependencies.
3. Set up the environment variables (under development).
   - For Windows:
     export .env vars in local dev using Git Bash:
    `
    export $(grep -v '^#' .env | xargs)
    `
4. Set up database tables and relationship.
   - For Windows:
     make sure you have installed XAMPP in your device and import the jobhun_tes.sql (do not forget to zip it first before importing)
4. Run the application using `go run main.go`.

## Support

If you find this project useful, please consider giving it a ⭐️ on [GitHub](https://github.com/gunturajip/golang-nonorm). Your support is greatly appreciated!

## Endpoints

### Mahasiswa

| Method | Endpoint                        | Description                  | Detail                                              |
| ------ | ------------------------------- | ---------------------------- | --------------------------------------------------- |
| POST   | `/api/v1/mahasiswa/`            | Create a new mahasiswa       | [Post Mahasiswa](#1-post-mahasiswa)                 |
| GET    | `/api/v1/mahasiswa/`            | Get a list of mahasiswa      | [Get All Mahasiswa](#2-get-all-mahasiswa)           |
| GET    | `/api/v1/mahasiswa/:id`         | Get a specific mahasiswa     | [Get Mahasiswa by ID](#3-get-mahasiswa-by-id)       |
| PUT    | `/api/v1/mahasiswa/:id`         | Update a specific mahasiswa  | [Update Mahasiswa by ID](#4-update-mahasiswa-by-id) |
| DELETE | `/api/v1/mahasiswa/:id`         | Delete a specific mahasiswa  | [Delete Mahasiswa by ID](#5-delete-mahasiswa-by-id) |

### Mahasiswa Endpoint Details

#### 1. Post Mahasiswa

**_Form data:_**

| Key                        | Value   | Description |
| -------------------------- | ------- | ----------- |
| nama                       | Kikik   |             |
| usia                       | 20      |             |
| gender                     | 0       |             |
| nama_jurusan               | Teknik  |             |
| nama_hobi                  | Tidur   |             |

**_Expected Response:_**

```json
{
    "status": 201,
    "success": "successfully created data mahasiswa (include jurusan and hobi)",
    "error": "",
    "data": {
        "mahasiswa": {
            "id": 6,
            "nama": "Kiko",
            "usia": 19,
            "gender": false,
            "tanggal_registrasi": "2023-04-23T23:19:42.2026943+07:00"
        },
        "jurusan": {
            "id": 6,
            "nama_jurusan": "Kimia"
        },
        "hobi": {
            "id": 4,
            "nama_hobi": "Gowes"
        },
        "mahasiswa_jurusan": {
            "id_mahasiswa": 6,
            "id_jurusan": 6
        },
        "mahasiswa_hobi": {
            "id_mahasiswa": 6,
            "id_hobi": 4
        }
    }
}
```

**_Status Code:_** 201

#### 2. Get All Mahasiswa

**_Expected Response:_**

```json
{
    "status": 200,
    "success": "successfully get all mahasiswa data including nama_jurusan and nama_hobi",
    "error": "",
    "data": [
        {
            "mahasiswa": {
                "id": 1,
                "nama": "Naila",
                "usia": 17,
                "gender": true,
                "tanggal_registrasi": "2023-04-23T12:16:39Z"
            },
            "nama_jurusan": "Keperawatan",
            "nama_hobi": "Makan"
        },
        {
            "mahasiswa": {
                "id": 2,
                "nama": "Hima",
                "usia": 8,
                "gender": true,
                "tanggal_registrasi": "2023-04-23T12:17:09Z"
            },
            "nama_jurusan": "Shinobi",
            "nama_hobi": "Makan"
        },
        {
            "mahasiswa": {
                "id": 3,
                "nama": "Alan",
                "usia": 13,
                "gender": false,
                "tanggal_registrasi": "2023-04-23T12:17:48Z"
            },
            "nama_jurusan": "Guru",
            "nama_hobi": "Minum"
        },
        {
            "mahasiswa": {
                "id": 4,
                "nama": "Guntur",
                "usia": 20,
                "gender": false,
                "tanggal_registrasi": "2023-04-23T12:51:38Z"
            },
            "nama_jurusan": "Programmer",
            "nama_hobi": "Makan"
        },
        {
            "mahasiswa": {
                "id": 5,
                "nama": "Kikik",
                "usia": 22,
                "gender": false,
                "tanggal_registrasi": "2023-04-23T13:11:03Z"
            },
            "nama_jurusan": "Teknisi",
            "nama_hobi": "Mukbang"
        },
        {
            "mahasiswa": {
                "id": 6,
                "nama": "Kiko",
                "usia": 19,
                "gender": false,
                "tanggal_registrasi": "2023-04-23T16:19:42Z"
            },
            "nama_jurusan": "Kimia",
            "nama_hobi": "Gowes"
        }
    ]
}
```

**_Status Code:_** 200

### 3. Get Mahasiswa by ID

### 4. Update Mahasiswa by ID

**_Form data:_**

| Key                        | Value   | Description |
| -------------------------- | ------- | ----------- |
| nama                       | Aji     |             |
| usia                       | 20      |             |
| gender                     | 0       |             |
| nama_jurusan               | Shinobi |             |
| nama_hobi                  | Makan   |             |

**_Expected Response:_**

```json
{
    "status": 200,
    "success": "successfully updated data mahasiswa (include jurusan and hobi)",
    "error": "",
    "data": {
        "mahasiswa": {
            "id": 2,
            "nama": "Aji",
            "usia": 20,
            "gender": false,
            "tanggal_registrasi": "0001-01-01T00:00:00Z"
        },
        "jurusan": {
            "id": 2,
            "nama_jurusan": "Shinobi"
        },
        "hobi": {
            "id": 1,
            "nama_hobi": "Makan"
        },
        "mahasiswa_jurusan": {
            "id_mahasiswa": 2,
            "id_jurusan": 2
        },
        "mahasiswa_hobi": {
            "id_mahasiswa": 2,
            "id_hobi": 1
        }
    }
}
```

**_Status Code:_** 200

### 5. Delete Mahasiswa by ID
