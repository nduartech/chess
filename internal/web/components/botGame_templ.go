// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"db"
	"engine"
	"github.com/olahol/melody"
	"layouts"
)

func NewBotGame(difficulty int, playAsWhite bool, m *melody.Melody, g *db.GameStateRow, d *db.DBConn) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			engine.ChessMatchWithBot(difficulty, playAsWhite, m, g, d)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row justify-center items-center p-4 m-20 bg-slate-600 hover:bg-slate-700 text-white text-md rounded\" id=\"newGameButton\" hx-trigger=\"click\" hx-get=\"/new-game\" hx-target=\"#start\" hx-swap=\"innerHTML\" hx-push-url=\"true\" hidden>New Game</div><div class=\"size-20rem md:size-40rem lg:size-50rem grid grid-cols-[1fr_1fr_1fr_1fr_1fr_1fr_1fr_1fr] grid-rows-[1fr_1fr_1fr_1fr_1fr_1fr_1fr_1fr] board\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if playAsWhite {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"A8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"H8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 2nd --> <div id=\"A7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"H7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><!-- 3th --> <div id=\"A6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"H6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 4st --> <div id=\"A5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"H5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><!-- 5th --> <div id=\"A4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"H4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 6th --> <div id=\"A3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"H3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><!-- id=\"3 7th --> <div id=\"A2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"H2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 8th --> <div id=\"A1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"H1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"H1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B1\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"A1\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 8th --> <div id=\"H2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B2\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"A2\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><!-- id=\"3 7th --> <div id=\"H3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B3\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"A3\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 6th --> <div id=\"H4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B4\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"A4\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><!-- 5th --> <div id=\"H5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B5\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"A5\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 4st --> <div id=\"H6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B6\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"A6\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><!-- 3th --> <div id=\"H7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"G7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"F7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"E7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"D7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"C7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"B7\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"A7\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><!-- 2nd --> <div id=\"H8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"G8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"F8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"E8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"D8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"C8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div><div id=\"B8\" class=\"black size-2.5rem md:size-5rem lg:size-6.25rem bg-[black] text-[5rem] text-center table-cell align-middle content-center square\"></div><div id=\"A8\" class=\"white size-2.5rem md:size-5rem lg:size-6.25rem text-[5rem] text-center table-cell align-middle bg-[white] content-center square\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"absolute top-0 left-1/12 right-1/12 left-1/12 w-10/12 h-11/12 bg-black flex flex-col justify-center items-center\" id=\"pawnPromotionContainer\" hidden><h1 class=\"text-white text-xl\">Pawn Promotion (Choose 1):</h1><div class=\"flex flex-col justify-center items-center w-fit h-fit space-y-2 px-2 py-2\" id=\"pawnPromotion\" data-move=\"\"><div class=\"flex flex-row justify-center items-center w-full text-xl text-white p-2 bg-slate-700 rounded\" data-piece=\"Q\">Queen</div><div class=\"flex flex-row justify-center items-center w-full text-xl text-white p-2 bg-slate-700 rounded\" data-piece=\"R\">Rook</div><div class=\"flex flex-row justify-center items-center w-full text-xl text-white p-2 bg-slate-700 rounded\" data-piece=\"B\">Bishop</div><div class=\"flex flex-row justify-center items-center w-full text-xl text-white p-2 bg-slate-700 rounded\" data-piece=\"N\">Knight</div></div></div><script type=\"module\">\n  import \"https://cdnjs.cloudflare.com/ajax/libs/dragula/3.7.3/dragula.min.js\";\n  const pieceMap = {\n    \"♖\": \"rookw\",\n    \"♘\": \"knightw\",\n    \"♗\": \"bishopw\",\n    \"♕\": \"queenw\",\n    \"♔\": \"kingw\",\n    \"♙\": \"pawnw\",\n    \"♟\": \"pawnb\",\n    \"♚\": \"kingb\",\n    \"♛\": \"queenb\",\n    \"♝\": \"bishopb\",\n    \"♞\": \"knightb\",\n    \"♜\": \"rookb\"\n  }\n\n  const getPieceHTML = (piece, num) => {\n    console.log(piece, num, pieceMap);\n    let pieceText = pieceMap[piece];\n    if (pieceText.startsWith(\"king\")) {\n      if (pieceText.endsWith(\"w\")) {\n        return `<div class=\"pwhite text-[sandybrown] select-none [float:initial] piece text-[2rem] md:text-[4rem] lg:text-[5rem] font-serif\" id=\"`+pieceText+`\">`+piece+`</div>`;\n      } else {\n        return `<div class=\"pblack text-[saddlebrown] select-none piece text-[2rem] md:text-[4rem] lg:text-[5rem] font-serif\" id=\"`+pieceText+`\">`+piece+`</div>`;\n      } \n    }\n    if (pieceText.endsWith(\"w\")) {\n      return `<div class=\"pwhite text-[sandybrown] select-none [float:initial] piece text-[2rem] md:text-[4rem] lg:text-[5rem] font-serif\" id=\"`+pieceText+num.toString()+`\">`+piece+`</div>`;\n    } else {\n      return `<div class=\"pblack text-[saddlebrown] select-none piece text-[2rem] md:text-[4rem] lg:text-[5rem] font-serif\" id=\"`+pieceText+num.toString()+`\">`+piece+`</div>`;\n    }\n  }\n\n  var drake = dragula([\n    // document.querySelector(\"#A1\"), document.querySelector(\"#B1\"), document.querySelector(\"#C1\"), document.querySelector(\"#D1\"), document.querySelector(\"#E1\"), document.querySelector(\"#F1\"), document.querySelector(\"#G1\"), document.querySelector(\"#H1\"),document.querySelector(\"#A2\"), document.querySelector(\"#B2\"), document.querySelector(\"#C2\"), document.querySelector(\"#D2\"), document.querySelector(\"#E2\"), document.querySelector(\"#F2\"), document.querySelector(\"#G2\"), document.querySelector(\"#H2\"),document.querySelector(\"#A3\"), document.querySelector(\"#B3\"), document.querySelector(\"#C3\"), document.querySelector(\"#D3\"), document.querySelector(\"#E3\"), document.querySelector(\"#F3\"), document.querySelector(\"#G3\"), document.querySelector(\"#H3\"),document.querySelector(\"#A4\"), document.querySelector(\"#B4\"), document.querySelector(\"#C4\"), document.querySelector(\"#D4\"), document.querySelector(\"#E4\"), document.querySelector(\"#F4\"), document.querySelector(\"#G4\"), document.querySelector(\"#H4\"),document.querySelector(\"#A5\"), document.querySelector(\"#B5\"), document.querySelector(\"#C5\"), document.querySelector(\"#D5\"), document.querySelector(\"#E5\"), document.querySelector(\"#F5\"), document.querySelector(\"#G5\"), document.querySelector(\"#H5\"),document.querySelector(\"#A6\"), document.querySelector(\"#B6\"), document.querySelector(\"#C6\"), document.querySelector(\"#D6\"), document.querySelector(\"#E6\"), document.querySelector(\"#F6\"), document.querySelector(\"#G6\"), document.querySelector(\"#H6\"),document.querySelector(\"#A7\"), document.querySelector(\"#B7\"), document.querySelector(\"#C7\"), document.querySelector(\"#D7\"), document.querySelector(\"#E7\"), document.querySelector(\"#F7\"), document.querySelector(\"#G7\"), document.querySelector(\"#H7\"),document.querySelector(\"#A8\"), document.querySelector(\"#B8\"), document.querySelector(\"#C8\"), document.querySelector(\"#D8\"), document.querySelector(\"#E8\"), document.querySelector(\"#F8\"), document.querySelector(\"#G8\"), document.querySelector(\"#H8\"),\n  ]);\n\n  // Function to animate the drag and drop\n  function animateDragAndDrop(el, source, target, pawnPromo) {\n    if (el.id.startsWith(\"king\")) {\n      if (source.id[0] === \"E\" && (target.id[0] === \"C\" || target.id[0] === \"G\")) {\n        // implementation of castling\n        let rook;\n        if (target.id[0] === \"C\") {\n          if (el.classList.contains(\"pblack\")) {\n            rook = document.getElementById(\"rookb1\");\n            rook.remove();\n            document.getElementById(\"D8\").appendChild(rook);\n          } else {\n            rook = document.getElementById(\"rookw1\");\n            rook.remove();\n            document.getElementById(\"D1\").appendChild(rook);\n          }\n        } else if (target.id[0] === \"G\") {\n          if (el.classList.contains(\"pblack\")) {\n            rook = document.getElementById(\"rookb2\");\n            rook.remove();\n            document.getElementById(\"F8\").appendChild(rook);\n          } else {\n            rook = document.getElementById(\"rookw2\");\n            rook.remove();\n            document.getElementById(\"F1\").appendChild(rook);\n          }\n        }\n      }\n    } else if (el.id.startsWith(\"pawn\")) {\n      if ((source.id[1] === \"2\" && target.id[1] === \"1\") || (source.id[1] === \"7\" && target.id[1] === \"8\")) {\n        // pawn promotion\n        let isBlackPiece = el.classList.contains(\"pblack\");\n        if (pawnPromo === \"Q\") {\n          el.id = el.id.replace(\"pawn\", \"queen\");\n          if (isBlackPiece) {\n            el.innerHTML = \"&#9819\";\n          } else {\n            el.innerHTML = \"&#9813\";\n          }\n        }\n        else if (pawnPromo === \"R\") {\n          el.id = el.id.replace(\"pawn\", \"rook\");\n          if (isBlackPiece) {\n            el.innerHTML = \"&#9820\";\n          } else {\n            el.innerHTML = \"&#9814\";\n          }\n        }\n        else if (pawnPromo === \"B\") {\n          el.id = el.id.replace(\"pawn\", \"bishop\");\n          if (isBlackPiece) {\n            el.innerHTML = \"&#9821\";\n          } else {\n            el.innerHTML = \"&#9815\";\n          }\n        }\n        else if (pawnPromo === \"N\") {\n          el.id = el.id.replace(\"pawn\", \"knight\");\n          if (isBlackPiece) {\n            el.innerHTML = \"&#9822\";\n          } else {\n            el.innerHTML = \"&#9817\";\n          }\n        }\n      } else if (source.id[0] !== target.id[0] && target.children.length === 0) {\n        document.getElementById(target.id[0] + source.id[1]).children[0].remove();\n      }\n    }\n    source.removeChild(el);\n    target.innerHTML = \"\";\n    target.appendChild(el);\n  }\n\n\n  let url = \"ws://\" + window.location.host + \"/ws\";\n  let ws = new WebSocket(url);\n  window.onbeforeunload = () => ws.close();\n  ws.onclose = function () {\n    alert(\"Game terminated in app, connection was closed, please refresh and try again\");\n  }\n  let containers = [];\n  let origContainers = [];\n  let targets = [];\n  let message;\n  ws.onmessage = (msg) => {\n    console.log(msg);\n    if (msg.data.startsWith(\"Init:\")) {\n      console.log(msg.data.slice(9,-1).split(\" \"));\n      const pieceTrack = new Map();\n      msg.data.slice(9,-1).split(\" \").forEach((piece) => {\n        if (pieceTrack.has(piece[3])) {\n          pieceTrack.set(piece[3], pieceTrack.get(piece[3])+1);\n        } else {\n          pieceTrack.set(piece[3], 1);\n        }\n        let pieceHTML = getPieceHTML(piece[3], pieceTrack.get(piece[3]));\n        document.getElementById(piece.slice(0,2).toUpperCase()).innerHTML = pieceHTML;\n      });\n    }\n    else if (msg.data.startsWith(\"Game completed.\")) {\n      alert(msg.data);\n      document.getElementById(\"newGameButton\").hidden = false;\n    } else if (msg.data.startsWith(\"Valid Moves:\")) {\n      message = msg;\n      msg.data.slice(13, -1).split(\" \").forEach((move) => {\n        drake.containers.push(document.getElementById(move.substring(0, 2).toUpperCase()));\n        containers.push(document.getElementById(move.substring(0, 2).toUpperCase()));\n        origContainers.push(document.getElementById(move.substring(0, 2).toUpperCase()));\n      });\n    } else {\n      console.log(msg.data);\n      if (msg.data.length === 4) {\n        animateDragAndDrop(document.getElementById(msg.data.substring(0, 2).toUpperCase()).children[0], document.getElementById(msg.data.substring(0, 2).toUpperCase()), document.getElementById(msg.data.substring(2, 4).toUpperCase()));\n      } else {\n        animateDragAndDrop(document.getElementById(msg.data.substring(0, 2).toUpperCase()).chidren[0], document.getElementById(msg.data.substring(0, 2).toUpperCase()), document.getElementById(msg.data.substring(2, 4).toUpperCase()), msg.data.substring(4).toUpperCase());\n      }\n    }\n  }\n  drake.on(\"drag\", (el, source) => {\n    for (let c of origContainers) {\n      if (c.id !== source.id) {\n        drake.containers.splice(c);\n        let ind = containers.indexOf(c);\n        containers.splice(ind, 1);\n      }\n    }\n    message.data.slice(13, -1).split(\" \").forEach((move) => {\n      console.log(move)\n      if (move.toUpperCase().startsWith(source.id)) {\n        drake.containers.push(document.getElementById(move.substring(2, 4).toUpperCase()));\n        targets.push(document.getElementById(move.substring(2, 4).toUpperCase()));\n      }\n    });\n    drake.containers.push(source);\n\n  });\n  drake.on(\"cancel\", function (el, container, source) {\n    targets.forEach((t) => drake.containers.splice(t));\n    containers.forEach((t) => drake.containers.splice(t));\n    targets = [];\n    containers = [...origContainers];\n    origContainers.forEach((ot) => drake.containers.push(ot));\n  });\n  drake.on(\"drop\", function (el, target, source, sibling) {\n    targets.forEach((t) => drake.containers.splice(t));\n    containers.forEach((t) => drake.containers.splice(t));\n    drake.containers.splice(source);\n    console.log(\"CONTAINERS\", drake.containers);\n    targets = [];\n    containers = [];\n    if (target.id !== source.id) {\n      console.log(\"DROPPED\", el, target, source, sibling);\n      let otherClass;\n      if (el.classList.contains(\"pwhite\")) {\n        otherClass = \"pblack\";\n      } else {\n        otherClass = \"pwhite\";\n      }\n      let pieceCaptured = \"\";\n      if (target.children.length > 1) {\n        if (target.children[0].classList.contains(otherClass)) {\n          target.removeChild(target.children[0]);\n        } else if (target.children[1].classList.contains(otherClass)) {\n          target.removeChild(target.children[1]);\n        }\n        pieceCaptured = \"x\";\n      }\n      let prefix = \"\";\n      if (el.id.startsWith(\"rook\")) {\n        prefix = \"R\";\n        ws.send(prefix + source.id.toLowerCase() + pieceCaptured + target.id.toLowerCase());\n      }\n      else if (el.id.startsWith(\"knight\")) {\n        prefix = \"N\";\n        ws.send(prefix + source.id.toLowerCase() + pieceCaptured + target.id.toLowerCase());\n      }\n      else if (el.id.startsWith(\"bishop\")) {\n        prefix = \"B\";\n        ws.send(prefix + source.id.toLowerCase() + pieceCaptured + target.id.toLowerCase());\n      }\n      else if (el.id.startsWith(\"queen\")) {\n        prefix = \"Q\";\n        ws.send(prefix + source.id.toLowerCase() + pieceCaptured + target.id.toLowerCase());\n      }\n      else if (el.id.startsWith(\"king\")) {\n        if (source.id[0] === \"E\" && (target.id[0] === \"C\" || target.id[0] === \"G\")) {\n          // implementation of castling\n          let rook;\n          if (target.id[0] === \"C\") {\n            if (el.classList.contains(\"pblack\")) {\n              rook = document.getElementById(\"rookb1\");\n              rook.remove();\n              document.getElementById(\"D8\").appendChild(rook);\n            } else {\n              rook = document.getElementById(\"rookw1\");\n              rook.remove();\n              document.getElementById(\"D1\").appendChild(rook);\n            }\n            ws.send(\"O-O-O\");\n          } else if (target.id[0] === \"G\") {\n            if (el.classList.contains(\"pblack\")) {\n              rook = document.getElementById(\"rookb2\");\n              rook.remove();\n              document.getElementById(\"F8\").appendChild(rook);\n            } else {\n              rook = document.getElementById(\"rookw2\");\n              rook.remove();\n              document.getElementById(\"F1\").appendChild(rook);\n            }\n            ws.send(\"O-O\");\n          }\n\n        } else {\n          prefix = \"K\";\n          ws.send(prefix + source.id.toLowerCase() + pieceCaptured + target.id.toLowerCase());\n        }\n      } else if (el.id.startsWith(\"pawn\")) {\n        if ((source.id[1] === \"2\" && target.id[1] === \"1\") || (source.id[1] === \"7\" && target.id[1] === \"8\")) {\n          // pawn promotion\n          document.getElementById(\"pawnPromotionContainer\").hidden = false;\n          if (pieceCaptured === \"x\") {\n            document.getElementById(\"pawnPromotion\").dataset.move = source.id.toLowerCase()[0] + pieceCaptured + target.id.toLowerCase();\n          } else {\n            document.getElementById(\"pawnPromotion\").dataset.move = target.id.toLowerCase();\n          }\n          Array.from(document.getElementById(\"pawnPromotion\").children).forEach((c) => {\n            c.addEventListener(\"click\", () => {\n              ws.send(document.getElementById(\"pawnPromotion\").dataset.move + \"=\" + c.dataset.piece);\n              document.getElementById(\"pawnPromotionContainer\").hidden = true;\n              let isBlackPiece = el.classList.contains(\"pblack\");\n              if (c.dataset.piece === \"Q\") {\n                el.id = el.id.replace(\"pawn\", \"queen\");\n                if (isBlackPiece) {\n                  el.innerHTML = \"&#9819\";\n                } else {\n                  el.innerHTML = \"&#9813\";\n                }\n              }\n              else if (c.dataset.piece === \"R\") {\n                el.id = el.id.replace(\"pawn\", \"rook\");\n                if (isBlackPiece) {\n                  el.innerHTML = \"&#9820\";\n                } else {\n                  el.innerHTML = \"&#9814\";\n                }\n              }\n              else if (c.dataset.piece === \"B\") {\n                el.id = el.id.replace(\"pawn\", \"bishop\");\n                if (isBlackPiece) {\n                  el.innerHTML = \"&#9821\";\n                } else {\n                  el.innerHTML = \"&#9815\";\n                }\n              }\n              else if (c.dataset.piece === \"N\") {\n                el.id = el.id.replace(\"pawn\", \"knight\");\n                if (isBlackPiece) {\n                  el.innerHTML = \"&#9822\";\n                } else {\n                  el.innerHTML = \"&#9817\";\n                }\n              }\n            }, {once: true});\n          });\n\n        } else {\n          // check if en passant, if so remove the captured pawn\n          if (source.id[0] !== target.id[0] && pieceCaptured === \"\") {\n            document.getElementById(target.id[0] + source.id[1]).children[0].remove();\n            pieceCaptured = \"x\";\n          }\n          ws.send(prefix + source.id.toLowerCase() + pieceCaptured + target.id.toLowerCase());\n        }\n      }\n\n    }\n  });\n\n\n        // Simulate Drag and Drop\n        // simulateDragAndDrop(drake,document.getElementById(\"pawnw1\"), document.getElementById(\"A3\"))\n</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.MainLayout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
