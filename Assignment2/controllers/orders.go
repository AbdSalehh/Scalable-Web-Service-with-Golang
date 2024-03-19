package controllers

import (
	"Assignment2/database"
	"Assignment2/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderData struct {
		OrderedAt    string            `json:"orderedAt"`
		CustomerName string            `json:"customerName"`
		Items        []models.ItemData `json:"items"`
	}

	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		http.Error(w, toJsonError("Invalid JSON data"), http.StatusBadRequest)
		return
	}

	order := models.Order{
		OrderedAt:    orderData.OrderedAt,
		CustomerName: orderData.CustomerName,
	}

	db := database.GetDB().Create(&order)
	if db.Error != nil {
		http.Error(w, toJsonError("Failed to create order"), http.StatusInternalServerError)
		return
	}

	var items []models.Item
	for _, itemData := range orderData.Items {
		item := models.Item{
			ItemCode:    itemData.ItemCode,
			Description: itemData.Description,
			Quantity:    itemData.Quantity,
			OrderID:     order.ID,
		}
		items = append(items, item)
	}

	if err := database.GetDB().Create(&items).Error; err != nil {
		http.Error(w, toJsonError("Failed to create items"), http.StatusInternalServerError)
		return
	}

	order.Items = items

	response := struct {
		Status  string       `json:"status"`
		Message string       `json:"message"`
		Order   models.Order `json:"order,omitempty"`
	}{
		Status:  "success",
		Message: "Created order successfully",
		Order:   order,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	db := database.GetDB().Preload("Items").Find(&orders)

	if db.Error != nil {
		http.Error(w, toJsonError("Failed to fetch orders"), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string          `json:"status"`
		Orders []models.Order `json:"orders,omitempty"`
	}{
		Status: "success",
		Orders: orders,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := mux.Vars(r)["orderId"]
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, toJsonError("Invalid order ID"), http.StatusBadRequest)
		return
	}

	var updateData struct {
		CustomerName string            `json:"customerName"`
		OrderedAt    string            `json:"orderedAt"`
		Items        []models.ItemData `json:"items"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, toJsonError("Invalid JSON data"), http.StatusBadRequest)
		return
	}

	var order models.Order
	db := database.GetDB().Preload("Items").First(&order, orderID)
	if db.Error != nil {
		http.Error(w, toJsonError("Order not found"), http.StatusNotFound)
		return
	}

	order.CustomerName = updateData.CustomerName
	order.OrderedAt = updateData.OrderedAt

	var itemUpdated bool
	for _, itemData := range updateData.Items {
		for i := range order.Items {
			if order.Items[i].ID == itemData.ID {
				order.Items[i].ItemCode = itemData.ItemCode
				order.Items[i].Description = itemData.Description
				order.Items[i].Quantity = itemData.Quantity

				// Update the item in the database
				if err := database.GetDB().Model(&order.Items[i]).Updates(order.Items[i]).Error; err != nil {
					http.Error(w, toJsonError("Failed to update item"), http.StatusInternalServerError)
					return
				}
				itemUpdated = true
				break
			}
		}
	}

	if !itemUpdated {
		http.Error(w, toJsonError("Failed to update order: line item not found"), http.StatusBadRequest)
		return
	}

	db = database.GetDB().Save(&order)
	if db.Error != nil {
		http.Error(w, toJsonError("Failed to update order"), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status  string       `json:"status"`
		Message string       `json:"message"`
		Order   models.Order `json:"order,omitempty"`
	}{
		Status:  "success",
		Message: "Updated order successfully",
		Order:   order,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := mux.Vars(r)["orderId"]
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, toJsonError("Invalid order ID"), http.StatusBadRequest)
		return
	}

	var order models.Order
	err = database.GetDB().Preload("Items").First(&order, orderID).Error
	if err != nil {
		http.Error(w, toJsonError("Order not found"), http.StatusNotFound)
		return
	}

	if len(order.Items) > 0 {
		err = database.GetDB().Delete(&order.Items).Error
		if err != nil {
			http.Error(w, toJsonError("Failed to delete items"), http.StatusInternalServerError)
			return
		}
	}

	err = database.GetDB().Delete(&order).Error
	if err != nil {
		http.Error(w, toJsonError("Failed to delete order"), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "success",
		Message: "Deleted order successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func toJsonError(message string) string {
	errorResponse := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "failed",
		Message: message,
	}
	jsonData, _ := json.Marshal(errorResponse)
	return string(jsonData)
}