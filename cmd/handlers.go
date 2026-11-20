package main

import "pizza-tracker/internal/models"

type Handler struct {
	orders *models.OrderModel;
	users *models.UserModel;
	notificationManager *NotificationManager;
}

func NewHandler(dbModel *models.DBModels) *Handler {
	return &Handler {
		orders: &dbModel.Order,
		users: &dbModel.User,
		notificationManager: NewNotificationManger(),
	}
}