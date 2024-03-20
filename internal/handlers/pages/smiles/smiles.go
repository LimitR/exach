package page_smiles

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Smiles(c *fiber.Ctx) error {
	smiles := ""
	for i := 1; i <= 43; i++ {
		smiles += `<img src="/public/img/smiles/` + strconv.Itoa(i) + `.gif" alt="" onclick="BB_code('f1', '[smile]` + strconv.Itoa(i) + `[/smile]', '')"> `
	}

	return c.Status(200).Send([]byte(`<input type="button" value="x" onclick="this.parentNode.innerHTML =''" style="float: right;"></button>` + smiles))
}
