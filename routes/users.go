package routes

import (
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
    // 創建一個空的用戶結構體
    user := models.User{}
    
    // 從 HTTP 請求的 JSON 數據中解析並綁定到用戶結構體中
    err := context.ShouldBindJSON(&user)
    
    // 如果解析失敗，返回錯誤響應
    if err != nil {
        context.JSON(400, gin.H{"message": "無法解析用戶數據"})
        return
    }
    
    // 嘗試將用戶信息保存到數據庫中
    err = user.Save()
    
    // 如果保存失敗，返回錯誤響應
    if err != nil {
        context.JSON(500, gin.H{"message": "無法創建用戶"})
        return
    }
    
    // 返回成功的響應
    context.JSON(201, gin.H{"message": "用戶創建成功！"})
}

func login(context *gin.Context) {
    var user models.User
    err := context.ShouldBindJSON(&user)

    if err != nil {
        context.JSON(400, gin.H{"message": "無法解析用戶數據"})
        return
    }

    err = user.ValidateCredentials()
    
    if err != nil {
        context.JSON(401, gin.H{"message": "登錄失敗"})
        return
    }

    context.JSON(200, gin.H{"message": "登錄成功！"})
}
