package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/term"
	"log"
	"strings"
	"time"
	"torctl/circuit"
	"torctl/collection"
	"torctl/config"
)

var options = []string{"New Circuit", "Get Status", "Help", "Exit"}

type Terminal struct {
	Auth         collection.Connection
	TerminalData []string
}

func (t *Terminal) New() {
	var temp string
	if t.Auth.Auth.AuthType == "password" {
		t.Auth.Conn, temp = config.PasswordAuth(t.Auth)
	} else if t.Auth.Auth.AuthType == "cookie" {
		t.Auth.Conn, temp = config.CookieAuth(t.Auth)
	} else {
		log.Fatal("Invalid Auth Type")
	}
	t.TerminalData = append(t.TerminalData, temp)
}

func (t *Terminal) StartUI() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()
	// Build Data return widget
	l := widgets.NewList()
	l.Title = "Data"
	l.TextStyle = termui.NewStyle(termui.ColorYellow)
	le, h, _ := term.GetSize(1)
	l.SetRect(le/2, 2, le-2, h-2)
	l.Rows = t.TerminalData
	// Build Options menu
	m := widgets.NewList()
	m.Title = "Options"
	m.TextStyle = termui.NewStyle(termui.ColorYellow)
	m.SetRect(2, 2, le/2, h/2-2)
	m.Rows = options
	// Build Status widget
	s := widgets.NewList()
	s.Title = "Status"
	s.TextStyle = termui.NewStyle(termui.ColorRed)
	s.SetRect(2, h/2, le/2, h-2)
	s.Rows = []string{"Status", "IPv4: " + t.Auth.Exits.V4, "IPv6: " + t.Auth.Exits.V6}

	c := widgets.NewParagraph()
	c.Title = "Torctl"
	c.SetRect(0, 0, le, h)
	c.BorderStyle.Bg = termui.ColorRed
	c.BorderStyle.Fg = termui.ColorWhite

	termui.Render(c, m, l, s)
	uiEvents := termui.PollEvents()

	data := circuit.GetConfig(t.Auth.Conn)
	t.TerminalData = append(t.TerminalData, data)
	l.Rows = t.TerminalData

	circuit.GetExits(&t.Auth)
	go func() {
		for true {
			le, h, _ := term.GetSize(1)
			m.SetRect(2, 2, le/2, h/2-2)
			l.SetRect(le/2, 2, le-2, h-2)
			s.SetRect(2, h/2, le/2, h-2)
			c.SetRect(0, 0, le, h)

			traffic := circuit.GetTraffic(t.Auth.Conn)
			s.Rows = []string{"Status", "IPv4: " + t.Auth.Exits.V4, "IPv6: " + t.Auth.Exits.V6, traffic}
			l.ScrollBottom()
			termui.Render(c, m, l, s)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			m.ScrollDown()
		case "k", "<Up>":
			m.ScrollUp()
		case "<Enter>":
			var response string
			l.Rows = append(l.Rows, "------"+options[m.SelectedRow]+"------")
			if options[m.SelectedRow] == "New Circuit" {
				response = circuit.UpdateCircuit(t.Auth)
				newExits := circuit.GetExits(&t.Auth)
				response = response + "\n" + newExits
			} else if options[m.SelectedRow] == "Help" {
				response = "PLACEHOLDER"
			} else if options[m.SelectedRow] == "Get Status" {
				response = circuit.GetStatus(t.Auth)
			} else if options[m.SelectedRow] == "Exit" {
				return
			}
			strings.Split(response, "\n")
			l.Rows = append(l.Rows, strings.Split(response+"\n", "\n")...)
		}
		termui.Render(m, l)
	}

}
