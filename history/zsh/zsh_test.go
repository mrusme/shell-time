package zsh

import (
	"strings"
	"testing"
)

func TestParseHistFile(t *testing.T) {
	testHistFile := `: 1624741139:0;sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \\
       https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'
: 1624745603:0;cargo install --locked code-minimap
: 1624745655:0;code-minimap
: 1624745816:0;vim .config/nvim/motd
: 1625525408:0;echo " â´´â´â´ Planning & Communication [Test]"
: 1627421927:0;curl -X "POST" "https://api.dev.test.dev/api" \\
     -H 'X-API-KEY: YqW06k7WA_NbhIT_ZtOcWlaa48MDySKR4d5gCcUvxs8=' \\
     -H 'Content-Type: application/json; charset=utf-8' \\
     -d $'{\
  "batch": {\
    "dateOfBirth": "1970-01-01",\
    "externalId1": "test",\
    "lastName": "Gonzales"\
  }\
}'
: 1629132862:0;/bin/ls -1 ./[0-9]+.pdf
: 1629831598:0;TESSDATA_PREFIX=~/projects/gitub/tessdata/ grim -g "2250,882 348x78" - | tesseract - - -l deu
: 1635604737:0;TF_VAR_fn_api_stripe=false TF_VAR_fn_api_blueprints=false TF_VAR_fn_api_jobs=false terraform apply`

	f := strings.NewReader(testHistFile)

	hist := History{}
	if err := hist.ParseHistFile(f); err != nil {
		t.Error(err)
	}

	// First line
	ts, cmd, args, err := hist.GetLine(0)
	if err != nil {
		t.Error(err)
	}

	if ts.Unix() != 1624741139 {
		t.Errorf("Parsing timestamp failed\n")
	}

	if cmd != "sh" {
		t.Errorf("Parsing command failed\n")
	}

	if args[1] != 'c' {
		t.Errorf("Parsing args failed\n")
	}

	// Second line
	ts, cmd, args, err = hist.GetLine(1)
	if err != nil {
		t.Error(err)
	}

	if ts.Unix() != 1624745603 {
		t.Errorf("Parsing timestamp failed\n")
	}

	if cmd != "cargo" {
		t.Errorf("Parsing command failed\n")
	}

	if args[0] != 'i' {
		t.Errorf("Parsing args failed\n")
	}

	// Sixth line
	ts, cmd, args, err = hist.GetLine(5)
	if err != nil {
		t.Error(err)
	}

	if ts.Unix() != 1627421927 {
		t.Errorf("Parsing timestamp failed\n")
	}

	if cmd != "curl" {
		t.Errorf("Parsing command failed\n")
	}

	if args[1] != 'X' {
		t.Errorf("Parsing args failed\n")
	}

	// Seventh line
	ts, cmd, args, err = hist.GetLine(6)
	if err != nil {
		t.Error(err)
	}

	if ts.Unix() != 1629132862 {
		t.Errorf("Parsing timestamp failed\n")
	}

	if cmd != "/bin/ls" {
		t.Errorf("Parsing command failed\n")
	}

	if args[1] != '1' {
		t.Errorf("Parsing args failed\n")
	}
}
