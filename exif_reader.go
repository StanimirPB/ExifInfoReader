package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
    "encoding/csv"
    "github.com/rwcarlsen/goexif/exif"
)

//==============================================================================
//settings
//var Dir = `source/` //относительно исходной дирректории
var Dir = `D:\Фото\` //абсолютный путь до фотографий
var outDir = `C:\Users\*******\go\src\parsers\exif\result\` //абсолютный путь к результату обработки
var outName = `res.csv` //название файла с результатами
var output, _ = os.Create(filepath.Join(outDir, filepath.Base(outName)))
//==============================================================================
var tags = []exif.FieldName {"ImageWidth",
                       "ImageLength",
                       "BitsPerSample",
                       "Compression",
                       "PhotometricInterpretation",
                       "Orientation",
                       "SamplesPerPixel",
                       "PlanarConfiguration",
                       "YCbCrSubSampling",
                       "YCbCrPositioning",
                       "XResolution",
                       "YResolution",
                       "ResolutionUnit",
                       "DateTime",
                       "ImageDescription",
                       "Make",
                       "Model",
                       "Software",
                       "Artist",
                       "Copyright",
                       "ExifIFDPointer",
                       "GPSInfoIFDPointer",
                       "InteroperabilityIFDPointer",
                       "ExifVersion",
                       "FlashpixVersion",
                       "ColorSpace",
                       "ComponentsConfiguration",
                       "CompressedBitsPerPixel",
                       "PixelXDimension",
                       "PixelYDimension",
                       "MakerNote",
                       "UserComment",
                       "RelatedSoundFile",
                       "DateTimeOriginal",
                       "DateTimeDigitized",
                       "SubSecTime",
                       "SubSecTimeOriginal",
                       "SubSecTimeDigitized",
                       "ImageUniqueID",
                       "ExposureTime",
                       "FNumber",
                       "ExposureProgram",
                       "SpectralSensitivity",
                       "ISOSpeedRatings",
                       "OECF",
                       "ShutterSpeedValue",
                       "ApertureValue",
                       "BrightnessValue",
                       "ExposureBiasValue",
                       "MaxApertureValue",
                       "SubjectDistance",
                       "MeteringMode",
                       "LightSource",
                       "Flash",
                       "FocalLength",
                       "SubjectArea",
                       "FlashEnergy",
                       "SpatialFrequencyResponse",
                       "FocalPlaneXResolution",
                       "FocalPlaneYResolution",
                       "FocalPlaneResolutionUnit",
                       "SubjectLocation",
                       "ExposureIndex",
                       "SensingMethod",
                       "FileSource",
                       "SceneType",
                       "CFAPattern",
                       "CustomRendered",
                       "ExposureMode",
                       "WhiteBalance",
                       "DigitalZoomRatio",
                       "FocalLengthIn35mmFilm",
                       "SceneCaptureType",
                       "GainControl",
                       "Contrast",
                       "Saturation",
                       "Sharpness",
                       "DeviceSettingDescription",
                       "SubjectDistanceRange",
                       "LensMake",
                       "LensModel",
                       "XPTitle",
	                     "XPComment",
	                     "XPAuthor",
	                     "XPKeywords",
	                     "XPSubject",
                       "ThumbJPEGInterchangeFormat",        // offset to thumb jpeg SOI
	                     "ThumbJPEGInterchangeFormatLength",  // byte length of thumb
                       "GPSVersionID",
  	                   "GPSLatitudeRef",
	                     "GPSLatitude",
	                     "GPSLongitudeRef",
	                     "GPSLongitude",
	                     "GPSAltitudeRef",
	                     "GPSAltitude",
	                     "GPSTimeStamp",
	                     "GPSSatelites",
	                     "GPSStatus",
	                     "GPSMeasureMode",
	                     "GPSDOP",
	                     "GPSSpeedRef",
	                     "GPSSpeed",
	                     "GPSTrackRef",
	                     "GPSTrack",
	                     "GPSImgDirectionRef",
	                     "GPSImgDirection",
	                     "GPSMapDatum",
	                     "GPSDestLatitudeRef",
	                     "GPSDestLatitude",
	                     "GPSDestLongitudeRef",
	                     "GPSDestLongitude",
	                     "GPSDestBearingRef",
	                     "GPSDestBearing",
	                     "GPSDestDistanceRef",
  	                   "GPSDestDistance",
	                     "GPSProcessingMethod",
	                     "GPSAreaInformation",
	                     "GPSDateStamp",
	                     "GPSDifferential",
                       "InteroperabilityIndex",
}

func main() {
  createHeadlers()

  exifCount := 0
  files := getFileslist()
  for k,file := range files {
    isExifCheck,metaData := getExif(file)
    if isExifCheck == true {exifCount++;go resaultSave(file.Name(),metaData)}
    fmt.Printf("\rОбработано файлов %v/%v <=> Exif %v",k+1,len(files),exifCount)
  }
}

func getFileslist() []os.FileInfo {
  files, err := ioutil.ReadDir(Dir)
  if err != nil {
      fmt.Printf("Ошибка чтения дирректории %v",err)
  }
  return files
}
func getExif(file os.FileInfo) (isExifCheck bool, metaData *exif.Exif) {
  isExifCheck = true
  imgFile,_ := os.Open(filepath.Join(Dir, filepath.Base(file.Name())))
  metaData,err := exif.Decode(imgFile)
  if err != nil {
    isExifCheck = false
  }
  return isExifCheck,metaData
}
func createHeadlers()  {
  var header []string
  header = append(header,"filename")
  header = append(header,"lat")
  header = append(header,"lng")
  for _,tag := range tags {
    header = append(header,fmt.Sprintf("%s", tag))
  }
  //заполняем csv
  writer := csv.NewWriter(output)
  writer.Comma = '|'
  defer writer.Flush()
  writer.Write(header)
}
func resaultSave(fName string,metaData *exif.Exif)  {
  var data []string
  data = append(data,fName) //file name
  lat,lng,_ := metaData.LatLong()
  data = append(data,fmt.Sprintf("%f", lat)) //dec latitude
  data = append(data,fmt.Sprintf("%f", lng)) //dec longitude
  for _,tag := range tags {
    md,_ := metaData.Get(tag)
    if md != nil {
      data = append(data,md.String())
    } else {
      data = append(data,"N/A")
    }
  }
  //заполняем csv
  writer := csv.NewWriter(output)
  writer.Comma = '|'
  defer writer.Flush()
  writer.Write(data)
}
