package models

/*
export interface menu {
header?: string;
title?: string;
icon?: object;
to?: string;
getURL?: boolean;
divider?: boolean;
chip?: string;
chipColor?: string;
chipVariant?: string;
chipIcon?: string;
children?: menu[];
disabled?: boolean;
type?: string;
subCaption?: string;
}
*/

type Menu struct {
	Header      string      `json:"header,omitempty"`
	Title       string      `json:"title,omitempty"`
	Icon        interface{} `json:"icon,omitempty"`
	To          string      `json:"to,omitempty"`
	GetURL      bool        `json:"get_url,omitempty"`
	Divider     bool        `json:"divider,omitempty"`
	Chip        string      `json:"chip,omitempty"`
	ChipColor   string      `json:"chip_color,omitempty"`
	ChipVariant string      `json:"chip_variant,omitempty"`
	ChipIcon    string      `json:"chip_icon,omitempty"`
	Children    []Menu      `json:"children,omitempty"`
	Disabled    bool        `json:"disabled,omitempty"`
	Type        string      `json:"type,omitempty"`
	SubCaption  string      `json:"sub_caption,omitempty"`
}
