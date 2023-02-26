package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

// Realación de uno a muchos
type Curse struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curse  Curse
}

func main() {

	//Relación de uno a uno
	/*
		juan := User{
			nombre: "juan",
			email:  "juan@gmail.com",
			activo: true,
		}

		david := User{
			nombre: "david",
			email:  "david@gmail.com",
			activo: true,
		}

		juanStudent := Student{
			user:   juan,
			codigo: "001rsd",
		}

		fmt.Println(david)
		fmt.Println(juanStudent.user.nombre)
	*/

	//Relación de uno a muchos
	video1 := Video{
		titulo: "01-Introducción",
	}

	video2 := Video{
		titulo: "02-Instalación",
	}

	curse := Curse{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curse = curse
	video2.curse = curse

	fmt.Println(video1.curse.titulo)

	for _, video := range curse.videos {
		fmt.Println(video.titulo)
	}
}
