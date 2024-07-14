package server

import (
    "strconv"
    "net/http"

    "github.com/gin-gonic/gin"
    "todoapp/todolist"
)

var list = todolist.NewList()

func GetList(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}

func AddItem(c *gin.Context) {
	content := c.PostForm("content")
	err := list.Add(content)
	var errString string
    if err == nil {
        errString = ""
    } else {
        errString = err.Error()
    }

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": errString,
	})
}

func SelectItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H{
		    "list": list.Items,
		    "error": err.Error(),
	    })
	}

	err = list.ToggleSelect(id)
	if err != nil {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H{
		    "list": list.Items,
		    "error": err.Error(),
	    })
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}

func DoneItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H{
		    "list": list.Items,
		    "error": err.Error(),
	    })
	}

	err = list.ToggleDone(id)
	if err != nil {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H{
		    "list": list.Items,
		    "error": err.Error(),
	    })
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}

func RemoveItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H{
		    "list": list.Items,
		    "error": err.Error(),
	    })
	}

	err = list.Remove(id)
	if err != nil {
	    c.HTML(http.StatusOK, "index.tmpl", gin.H{
		    "list": list.Items,
		    "error": err.Error(),
	    })
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}

func RemoveAll(c *gin.Context) {
    list.RemoveAll()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}

func RemoveSelecteds(c *gin.Context) {
    list.RemoveSelecteds()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}

func DoneSelecteds(c *gin.Context) {
    list.ToggleDoneSelecteds()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": list.Items,
		"error": "",
	})
}
