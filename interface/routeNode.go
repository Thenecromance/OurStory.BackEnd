//package Interface
//
//import (
//	"encoding/json"
//	"github.com/Thenecromance/OurStories/base/log"
//	"github.com/gin-gonic/gin"
//)
//
//type RouteNode struct {
//	unHandled        map[string]*RouteNode `json:"-"`
//	ParentPtr        *RouteNode            `json:"-"` //disable json marshalling and unmarshalling or else it will cause a stack overflow
//	*gin.RouterGroup `json:"-"`
//
//	Name     string                `json:"node"`
//	Parent   string                `json:"parent"`
//	Children map[string]*RouteNode `json:"children"`
//}
//
//func (rn *RouteNode) IsRoot() bool {
//	return rn.ParentPtr == nil
//}
//
//func (rn *RouteNode) SetParent(parent *RouteNode) {
//
//	rn.ParentPtr = parent
//	parent.Children[rn.Name] = rn
//}
//
//func (rn *RouteNode) LoadController(node ...*RouteNode) {
//	if rn.unHandled == nil {
//		rn.unHandled = make(map[string]*RouteNode)
//	}
//	for _, n := range node {
//		rn.unHandled[n.Name] = n
//	}
//}
//
//func (rn *RouteNode) MakeAsTree() error {
//	////dispatch all the unhandled nodes to their parent
//	//for _, node := range rn.unHandled {
//	//	parentNode := rn.unHandled[node.Parent] // get the parent node from the unhandled map
//	//	if parentNode == nil {
//	//		if node.Parent == "/" {
//	//			node.SetParent(rn)
//	//		} else {
//	//			log.Error(fmt.Sprintf("Parent node %s not found for node %s", node.Parent, node.Name))
//	//			return errors.New(fmt.Sprintf("Parent node %s not found for node %s", node.Parent, node.Name))
//	//		}
//	//	} else {
//	//		node.SetParent(parentNode)
//	//	}
//	//}
//	//rn.unHandled = nil
//	//return nil
//
//	//iterate through all the unhandled nodes
//	for _, node := range rn.unHandled {
//		//if the node set the parent node's name is this , then set the parent node to this
//		if node.Parent == rn.Name {
//			node.SetParent(rn)
//		} else {
//			//otherwise, search other unhandled nodes for the parent node
//			parentNode := rn.unHandled[node.Parent]
//			if parentNode != nil {
//				node.SetParent(parentNode)
//			} else {
//				log.Errorf("Parent node %s not found for node %s", node.Parent, node.Name)
//			}
//		}
//
//	}
//	return nil
//}
//
//func (rn *RouteNode) CreateNodeGroups() {
//	log.Info("Building routes for ", rn.Path())
//	for _, node := range rn.Children {
//		node.RouterGroup = rn.Group(node.Name)
//		node.CreateNodeGroups()
//	}
//}
//
//func (rn *RouteNode) path() string {
//	if rn.IsRoot() {
//		return rn.Name
//	}
//	return rn.ParentPtr.path() + "/" + rn.Name
//}
//func (rn *RouteNode) Path() string {
//	p := rn.path()
//	if p == "/" {
//		return p
//	}
//	return p[1:]
//}
//
//// String will return the json string of this RouteNode
//func (rn *RouteNode) String() string {
//	marshal, err := json.MarshalIndent(rn, "", "    ")
//	if err != nil {
//		return ""
//	}
//	return string(marshal)
//}
//
//func NewRootNode() *RouteNode {
//	return &RouteNode{
//		Name:     "/",
//		Parent:   "",
//		Children: make(map[string]*RouteNode),
//	}
//}
//func NewNode(parent, name string) *RouteNode {
//	return &RouteNode{
//		Name:     name,
//		Parent:   parent,
//		Children: make(map[string]*RouteNode),
//	}
//}

package Interface

import "github.com/gin-gonic/gin"

type GroupNode struct {
	Parent     string   `json:"parent"`
	MiddleWare []string `json:"middleWare"`

	Router *gin.RouterGroup `json:"-"`
}
