/*
节点
*/

package gopher

import (
	"net/http"
)

// URL: /nodes
// 列出所有节点及其详细信息
func nodesHandler(w http.ResponseWriter, r *http.Request) {
	var nodes []Node

	c := DB.C("nodes")
	c.Find(nil).Sort("-topiccount").All(&nodes)

	renderTemplate(w, r, "node/list.html", map[string]interface{}{"nodes": nodes})
}
