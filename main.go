package main

import (
    "fmt"
    "gocv.io/x/gocv"
)

func main() {
    // Открываем видеопоток с камеры
    webcam, err := gocv.OpenVideoCapture(0)
    if err != nil {
        fmt.Println("Ошибка открытия видеопотока:", err)
        return
    }
    defer webcam.Close()

    // Создаем окно для отображения видеопотока
    window := gocv.NewWindow("Webcam")
    defer window.Close()

    img := gocv.NewMat()
    defer img.Close()

    for {
        if ok := webcam.Read(&img); !ok {
            fmt.Println("Не удалось прочитать кадр из видеопотока")
            return
        }
        if img.Empty() {
            continue
        }

        // Отображаем кадр в окне
        window.IMShow(img)
        if window.WaitKey(1) >= 0 {
            break
        }
    }
}