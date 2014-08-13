package gosseract_test

import "github.com/otiai10/gosseract"
import "testing"
import "fmt"
import "image"

func ExampleMust(t *testing.T) {
	// TODO: it panics! error handling in *Client.accept
	out := gosseract.Must(map[string]string{"src": "./.samples/png/sample002.png"})
	fmt.Println(out)
}

func ExampleClient_Src(t *testing.T) {
	client, _ := gosseract.NewClient()
	out, _ := client.Src("./samples/png/samples000.png").Out()
	fmt.Println(out)
}

func ExampleClient_Image(t *testing.T) {
	client, _ := gosseract.NewClient()
	var img image.Image // any your image instance
	out, _ := client.Image(img).Out()
	fmt.Println(out)
}
