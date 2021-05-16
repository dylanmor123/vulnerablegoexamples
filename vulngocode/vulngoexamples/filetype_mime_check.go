
		switch filetype {
		case "image/jpeg", "image/jpg":
			fmt.Println(filetype)
		case "image/gif":
			fmt.Println(filetype)
		case "image/png":
			fmt.Println(filetype)
		default:
			fmt.Println("unknown file type uploaded")
		}
