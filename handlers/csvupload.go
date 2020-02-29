package handlers

// CSVUpload upload with csv
// func CSVUpload(c echo.Context) error {
// 	defer c.Request().Body.Close()

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	src, err := file.Open()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	defer src.Close()

// 	// Destination
// 	dst, err := os.Create("assets/files/file.csv")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	defer dst.Close()

// 	// Copy
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	// Open the file
// 	csvfile, err := os.Open("assets/files/file.csv")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	// Parse the file
// 	r := csv.NewReader(csvfile)
// 	//r := csv.NewReader(bufio.NewReader(csvfile))
// 	i := 0
// 	// Iterate through the records
// 	for {
// 		if i%50 == 0 {
// 			time.Sleep(2 * time.Second)
// 		}
// 		// Read each record from csv
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			return c.JSON(http.StatusOK, map[string]interface{}{
// 				"message": "Success",
// 			})
// 		}

// 		type Filter struct {
// 			Name string `json:"name"`
// 		}

// 		//provinsi
// 		prov := models.Provinsi{}
// 		err = prov.SingleFindFilter(&Filter{
// 			Name: record[7],
// 		})
// 		if err != nil {
// 			prov.Name = record[7]
// 			prov.Create()
// 		}

// 		//Kota
// 		city := models.Kota{}
// 		err = city.SingleFindFilter(&Filter{
// 			Name: record[6],
// 		})
// 		if err != nil {
// 			city.Name = record[6]
// 			city.ProvinsiID = prov.BaseModel.ID
// 			city.Type = record[5]
// 			city.Create()
// 		}

// 		//Kecamatan
// 		district := models.Kecamatan{}
// 		err = district.SingleFindFilter(&Filter{
// 			Name: record[4],
// 		})
// 		if err != nil {
// 			district.Name = record[4]
// 			district.KotaID = city.BaseModel.ID
// 			district.Create()
// 		}

// 		//Kelurahan
// 		village := models.Kelurahan{}
// 		err = village.SingleFindFilter(&Filter{
// 			Name: record[2],
// 		})
// 		if err != nil {
// 			village.Name = record[2]
// 			village.PostalCode = record[1]
// 			village.AreaCode = record[3]
// 			village.KecamatanID = district.BaseModel.ID
// 			village.Create()
// 		}
// 		i++
// 	}

// 	// return c.JSON(http.StatusOK, map[string]interface{}{
// 	// 	"message": "Data Berhasil Di input",
// 	// })
// }
