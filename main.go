package main

import (
	"copy_image/images"
	"flag"
)

func main() {
	var imageText string
	var localImageDir string
	var harborAddr string
	var harborUser string
	var harborPassword string
	var operator string

	flag.StringVar(&imageText, "t", "", "The image list that needs to be pulled")
	flag.StringVar(&localImageDir, "d", "", "The address where the image manifests is stored this time")
	flag.StringVar(&harborAddr, "hd", "", "harbor address")
	flag.StringVar(&harborUser, "u", "", "user for harbor")
	flag.StringVar(&harborPassword, "p", "", "password for harbor")
	flag.StringVar(&operator, "o", "pull", "pull or push")
	flag.Parse()

	images.Images(imageText, localImageDir, harborAddr, operator, harborUser, harborPassword)
}

/*
go build -tags=containers_image_openpgp -o pullImage main.go

./pullImage -t ./1.txt -d ./temp_manifests -o pull
./pullImage -t ./1.txt -d ./temp_manifests -o pull -u  -p

[root@k8smaster go]# docker load -i library_golang_1_24_0.tar
01c9a2a5f237: Loading layer  121.3MB/121.3MB
f8217d7865d2: Loading layer  49.58MB/49.58MB
20a9b386e10e: Loading layer  181.9MB/181.9MB
bb937a58bbf5: Loading layer  261.2MB/261.2MB
fde374622a0b: Loading layer  263.8MB/263.8MB
95e67e04fbca: Loading layer   2.56kB/2.56kB
5f70bf18a086: Loading layer  1.024kB/1.024kB
Loaded image ID: sha256:245780beb82f97496ad0dee2d70c6fd2bfa56adedb229a3e25c9e801dd262b5e
*/
