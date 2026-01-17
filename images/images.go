package images

import (
	"bufio"
	"context"
	"fmt"
	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/transports/alltransports"
	"os"
	"strings"
)

// sanitizeImageName 清理镜像名称以适合作为文件名
func sanitizeImageName(imageName string) string {
	// 替换不允许的字符
	replacer := strings.NewReplacer(
		"/", "_",
		":", "_",
		".", "_",
		"+", "_",
	)
	return replacer.Replace(imageName)
}

func Images(imageText, localImageDir, harborAddr, operator, harborUser, harborPasswd string) {
	policyContext, err := getPolicyContext()
	if err != nil {
		fmt.Println(err)
	}
	defer policyContext.Destroy()

	file, err := os.Open(imageText)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
		return
	}

	modified := false
	for i, line := range lines {
		oriLine := line

		var srcImage string
		var destImage string
		switch operator {
		case "pull":
			srcImage = fmt.Sprintf("docker://%v", line)
			destImage = fmt.Sprintf("docker-archive:%v/%v.tar", localImageDir, sanitizeImageName(line))
		}
		fmt.Printf("srcImage:%+v\n", srcImage)
		fmt.Printf("destImage:%+v\n", destImage)

		srcRef, err := alltransports.ParseImageName(srcImage)
		if err != nil {
			fmt.Println(err)
			continue
		}
		destRef, err := alltransports.ParseImageName(destImage)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("srcRef:%+v\n", srcRef)
		fmt.Printf("destRef:%+v\n", destRef)

		srcContext := systemContext(harborUser, harborPasswd)
		destContext := systemContext(harborUser, harborPasswd)

		_, err = copy.Image(context.Background(), policyContext, destRef, srcRef, &copy.Options{
			ReportWriter:   os.Stdout,
			SourceCtx:      srcContext,
			DestinationCtx: destContext,
		})
		if err != nil {
			fmt.Println(err)
			lines[i] = oriLine + "  fail"
			continue
		}

		lines[i] = oriLine + "  success"
		modified = true
		fmt.Printf("[%v] is successfully copied\n", oriLine)
	}

	if modified {
		if err := os.WriteFile(imageText, []byte(strings.Join(lines, "\n")), 0644); err != nil {
			fmt.Println("文件更新失败:", err)
			return
		}
	}
}
