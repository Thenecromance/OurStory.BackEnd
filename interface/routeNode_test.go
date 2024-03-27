package Interface

import (
	"reflect"
	"testing"
)

/*
func TestNewNode(t *testing.T) {
	type args struct {
		parent string
		name   string
	}
	tests := []struct {
		name string
		args args
		want *RouteNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.parent, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestNewRootNode(t *testing.T) {
	tests := []struct {
		name string
		want *RouteNode
	}{
		// TODO: Add test cases.
		{
			name: "TestNewRootNode",
			want: &RouteNode{
				Name:     "/",
				Parent:   "",
				Children: make(map[string]*RouteNode),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRootNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRootNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteNode_IsRoot(t *testing.T) {
	type fields struct {
		unHandled map[string]*RouteNode
		ParentPtr *RouteNode
		Name      string
		Parent    string
		Children  map[string]*RouteNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &RouteNode{
				unHandled: tt.fields.unHandled,
				ParentPtr: tt.fields.ParentPtr,
				Name:      tt.fields.Name,
				Parent:    tt.fields.Parent,
				Children:  tt.fields.Children,
			}
			if got := rn.IsRoot(); got != tt.want {
				t.Errorf("IsRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteNode_Load(t *testing.T) {
	type fields struct {
		unHandled map[string]*RouteNode
		ParentPtr *RouteNode
		Name      string
		Parent    string
		Children  map[string]*RouteNode
	}
	type args struct {
		node []*RouteNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &RouteNode{
				unHandled: tt.fields.unHandled,
				ParentPtr: tt.fields.ParentPtr,
				Name:      tt.fields.Name,
				Parent:    tt.fields.Parent,
				Children:  tt.fields.Children,
			}
			rn.Load(tt.args.node...)
		})
	}
}

func TestRouteNode_MakeAsTree(t *testing.T) {
	type fields struct {
		unHandled map[string]*RouteNode
		ParentPtr *RouteNode
		Name      string
		Parent    string
		Children  map[string]*RouteNode
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &RouteNode{
				unHandled: tt.fields.unHandled,
				ParentPtr: tt.fields.ParentPtr,
				Name:      tt.fields.Name,
				Parent:    tt.fields.Parent,
				Children:  tt.fields.Children,
			}
			if err := rn.MakeAsTree(); (err != nil) != tt.wantErr {
				t.Errorf("MakeAsTree() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteNode_Path(t *testing.T) {
	type fields struct {
		unHandled map[string]*RouteNode
		ParentPtr *RouteNode
		Name      string
		Parent    string
		Children  map[string]*RouteNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &RouteNode{
				unHandled: tt.fields.unHandled,
				ParentPtr: tt.fields.ParentPtr,
				Name:      tt.fields.Name,
				Parent:    tt.fields.Parent,
				Children:  tt.fields.Children,
			}
			if got := rn.Path(); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteNode_SetParent(t *testing.T) {
	type fields struct {
		unHandled map[string]*RouteNode
		ParentPtr *RouteNode
		Name      string
		Parent    string
		Children  map[string]*RouteNode
	}
	type args struct {
		parent *RouteNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &RouteNode{
				unHandled: tt.fields.unHandled,
				ParentPtr: tt.fields.ParentPtr,
				Name:      tt.fields.Name,
				Parent:    tt.fields.Parent,
				Children:  tt.fields.Children,
			}
			rn.SetParent(tt.args.parent)
		})
	}
}
