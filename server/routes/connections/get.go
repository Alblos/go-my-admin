package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/types"
	routeTypes "github.com/go-my-admin/server/routes/types"
	"strconv"
)

type HandleGetAllConectionsResponse struct {
	Error       bool               `json:"error"`
	Connections []types.Connection `json:"connections"`
}

// HandleGetAllConnections
// @BasePath /
// @Summary Get all connections
// @Description Gets all the connections that a user has added
// @Tags connections
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} HandleGetAllConectionsResponse "Returns all connections"
// @Failure 400 {object} types.ErrorResponse "Invalid connection ID"
// @Failure 404 {object} types.ErrorResponse "Connection with the given ID does not exist or you don't have permission to view it"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /connections [get]
func HandleGetAllConnections(c *gin.Context) {
	db := database.InternalDb

	connectionsInDb, err := db.RunQueryWithParams("SELECT id, common_name, database_name, host, port, username, ssl_mode FROM connections WHERE owner_id = (SELECT id FROM users WHERE username = $1)", c.GetString("username"))
	if err != nil {
		c.JSON(500, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Error getting connections: " + err.Error(),
		})
		return
	}

	var connectionsArray []types.Connection

	for connectionsInDb.Next() {
		var connection types.Connection
		err = connectionsInDb.Scan(&connection.Id, &connection.CommonName, &connection.DatabaseName, &connection.Host, &connection.Port, &connection.Username, &connection.SslMode)
		if err != nil {
			c.JSON(500, routeTypes.ErrorResponse{
				Error:   true,
				Message: "Error scanning connection: " + err.Error(),
			})
			return
		}
		connectionsArray = append(connectionsArray, connection)
	}

	c.JSON(200, HandleGetAllConectionsResponse{
		Error:       false,
		Connections: connectionsArray,
	})
}

type HandleGetConnectionByIdResponse struct {
	Error      bool             `json:"error"`
	Connection types.Connection `json:"connection"`
}

// HandleGetConnectionById
// @BasePath /
// @Summary Get a connection by ID
// @Description Get the data of a connection by its ID
// @Tags connections
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Connection ID"
// @Success 200 {object} HandleGetConnectionByIdResponse "Returns the connection data"
// @Failure 400 {object} types.ErrorResponse "Invalid connection ID"
// @Failure 404 {object} types.ErrorResponse "Connection with the given ID does not exist"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /connections/{id} [get]
func HandleGetConnectionById(c *gin.Context) {
	db := database.InternalDb

	connectionId := c.Param("id")

	id, err := strconv.Atoi(connectionId)
	if err != nil {
		c.JSON(400, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Invalid connection ID",
		})
		return
	}
	exists, err := connections.CheckIfConnectionWithIdExists(id)
	if err != nil {
		c.JSON(500, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Error checking if connection exists: " + err.Error(),
		})
		return
	}
	if !exists {
		c.JSON(404, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Connection with the given ID does not exist",
		})
		return
	}

	connection, err := db.RunQueryWithParams("SELECT c.id, common_name, database_name, host, port, c.username, ssl_mode, u.username FROM connections c INNER JOIN public.users u on u.id = c.owner_id WHERE c.id = $1", id)
	if err != nil {
		c.JSON(500, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Error getting connection: " + err.Error(),
		})
		return
	}

	var connectionMapped types.Connection

	for connection.Next() {
		var username string
		err = connection.Scan(&connectionMapped.Id, &connectionMapped.CommonName, &connectionMapped.DatabaseName, &connectionMapped.Host, &connectionMapped.Port, &connectionMapped.Username, &connectionMapped.SslMode, &username)
		if err != nil {
			c.JSON(500, routeTypes.ErrorResponse{
				Error:   true,
				Message: "Error scanning connection: " + err.Error(),
			})
			return
		}
		if username != c.GetString("username") {
			c.JSON(404, routeTypes.ErrorResponse{
				Error:   true,
				Message: "Connection with the given ID does not exist",
			})
			return
		}
	}

	c.JSON(200, HandleGetConnectionByIdResponse{
		Error:      false,
		Connection: connectionMapped,
	})
}
