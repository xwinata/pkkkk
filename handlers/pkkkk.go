package handlers

import (
	"net/http"
	"pkkkk/appconf"
	"pkkkk/models"

	"github.com/labstack/echo"
)

// Result shows query result
type Result struct {
	Provinsi  models.Provinsi  `json:"provinsi"`
	Kota      models.Kota      `json:"kota"`
	Kecamatan models.Kecamatan `json:"kecamatan"`
	Kelurahan models.Kelurahan `json:"kelurahan"`
	Options   interface{}      `json:"options"`
}

// Start by showing all provinsi
func Start(c echo.Context) (err error) {
	defer c.Request().Body.Close()

	provinsi := []models.Provinsi{}

	err = appconf.Configuration.DB.Find(&provinsi).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, provinsi)
}

// // InProvinsi shows provinsi content
// func InProvinsi(c echo.Context) error {
// 	defer c.Request().Body.Close()

// 	kota := models.Kota{}
// 	idprovinsi, err := strconv.ParseUint(c.Param("idprovinsi"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	provinsi := models.Provinsi{}
// 	provinsi.FindbyID(idprovinsi)

// 	options, err := kota.FindFilter([]string{"name"}, []string{"asc"}, 0, 0, &models.Kota{
// 		ProvinsiID: idprovinsi,
// 	})
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, &Result{
// 		Provinsi: provinsi,
// 		Options:  options,
// 	})
// }

// // InKota shows kota content
// func InKota(c echo.Context) error {
// 	defer c.Request().Body.Close()

// 	kecamatan := models.Kecamatan{}
// 	idprovinsi, err := strconv.ParseUint(c.Param("idprovinsi"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	idkota, err := strconv.ParseUint(c.Param("idkota"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	provinsi := models.Provinsi{}
// 	kota := models.Kota{}
// 	provinsi.FindbyID(idprovinsi).Related(&kota)

// 	options, err := kecamatan.FindFilter([]string{"name"}, []string{"asc"}, 0, 0, &models.Kecamatan{
// 		KotaID: idkota,
// 	})
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, &Result{
// 		Provinsi: provinsi,
// 		Kota:     kota,
// 		Options:  options,
// 	})
// }

// // InKecamatan shows kecamatan content
// func InKecamatan(c echo.Context) error {
// 	defer c.Request().Body.Close()

// 	kelurahan := models.Kelurahan{}
// 	idprovinsi, err := strconv.ParseUint(c.Param("idprovinsi"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	idkota, err := strconv.ParseUint(c.Param("idkota"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	idkecamatan, err := strconv.ParseUint(c.Param("idkecamatan"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	provinsi := models.Provinsi{}
// 	provinsi.FindbyID(idprovinsi)
// 	kota := models.Kota{}
// 	kota.FindbyID(idkota)
// 	kecamatan := models.Kecamatan{}
// 	kecamatan.FindbyID(idkecamatan)

// 	options, err := kelurahan.FindFilter([]string{"name"}, []string{"asc"}, 0, 0, &models.Kelurahan{
// 		KecamatanID: idkecamatan,
// 	})
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, &Result{
// 		Provinsi:  provinsi,
// 		Kota:      kota,
// 		Kecamatan: kecamatan,
// 		Options:   options,
// 	})
// }

// // InKelurahan shows kelurahan content
// func InKelurahan(c echo.Context) error {
// 	defer c.Request().Body.Close()

// 	idprovinsi, err := strconv.ParseUint(c.Param("idprovinsi"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	idkota, err := strconv.ParseUint(c.Param("idkota"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	idkecamatan, err := strconv.ParseUint(c.Param("idkecamatan"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}
// 	idkelurahan, err := strconv.ParseUint(c.Param("idkelurahan"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
// 			"message": fmt.Sprintf("%v", err),
// 		})
// 	}

// 	provinsi := models.Provinsi{}
// 	provinsi.FindbyID(idprovinsi)
// 	kota := models.Kota{}
// 	kota.FindbyID(idkota)
// 	kecamatan := models.Kecamatan{}
// 	kecamatan.FindbyID(idkecamatan)
// 	kelurahan := models.Kelurahan{}
// 	kelurahan.FindbyID(idkelurahan)

// 	return c.JSON(http.StatusOK, &Result{
// 		Provinsi:  provinsi,
// 		Kota:      kota,
// 		Kecamatan: kecamatan,
// 		Kelurahan: kelurahan,
// 	})
// }
