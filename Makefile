build:
	go build

# Haven't tested this yet
run:
	ifneq ("$(wildcard $(t2t_image_service.exe))","")
		echo "In windows"
	else
		echo "NOt windwos"
	endif
