package app

import (
	"context"
	"net/http"

	"github.com/acoshift/acourse/pkg/acourse"
	"github.com/acoshift/httperror"
	"google.golang.org/grpc/metadata"
	"gopkg.in/gin-gonic/gin.v1"
)

func makeServiceContext(r *http.Request) context.Context {
	md := metadata.MD{}
	header := r.Header
	if v := header.Get("Authorization"); v != "" {
		md = metadata.Join(md, metadata.Pairs("authorization", v))
	}
	return metadata.NewContext(context.Background(), md)
}

// RegisterUserServiceClient registers a User service client to http server
func RegisterUserServiceClient(httpServer *gin.Engine, s acourse.UserServiceClient) {
	httpServer.POST("/acourse.UserService/GetUser", func(ctx *gin.Context) {
		req := new(acourse.GetUserRequest)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		res, err := s.GetUser(makeServiceContext(ctx.Request), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleOK(ctx, res)
	})

	httpServer.GET("/acourse.UserService/GetMe", func(ctx *gin.Context) {
		res, err := s.GetMe(makeServiceContext(ctx.Request), new(acourse.Empty))
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleOK(ctx, res)
	})

	httpServer.POST("/acourse.UserService/UpdateMe", func(ctx *gin.Context) {
		req := new(acourse.User)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		_, err = s.UpdateMe(makeServiceContext(ctx.Request), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleSuccess(ctx)
	})
}

// RegisterEmailServiceClient registers a Email service client to http server
func RegisterEmailServiceClient(httpServer *gin.Engine, s acourse.EmailServiceClient) {
	httpServer.POST("/acourse.EmailService/Send", func(ctx *gin.Context) {
		req := new(acourse.Email)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		_, err = s.Send(makeServiceContext(ctx.Request), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleSuccess(ctx)
	})
}

// RegisterPaymentService registers a payment service
func RegisterPaymentService(service *gin.Engine, s PaymentService) {
	service.POST("/acourse.PaymentService/ListPayments", func(ctx *gin.Context) {
		req := new(PaymentListRequest)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		res, err := s.ListPayments(ctx.Request.Context(), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleOK(ctx, res.Expose())
	})

	service.POST("/acourse.PaymentService/ApprovePayments", func(ctx *gin.Context) {
		req := new(IDsRequest)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		err = s.ApprovePayments(ctx.Request.Context(), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleSuccess(ctx)
	})

	service.POST("/acourse.PaymentService/RejectPayments", func(ctx *gin.Context) {
		req := new(IDsRequest)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		err = s.RejectPayments(ctx.Request.Context(), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleSuccess(ctx)
	})
}

// RegisterCourseService registers a course service
func RegisterCourseService(service *gin.Engine, s CourseService) {
	service.POST("/acourse.CourseService/ListCourses", func(ctx *gin.Context) {
		req := new(CourseListRequest)
		err := ctx.BindJSON(req)
		if err != nil {
			handleError(ctx, httperror.BadRequestWith(err))
			return
		}
		res, err := s.ListCourses(ctx.Request.Context(), req)
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleOK(ctx, res.Expose())
	})

	service.GET("/acourse.CourseService/ListEnrolledCourses", func(ctx *gin.Context) {
		res, err := s.ListEnrolledCourses(ctx.Request.Context())
		if err != nil {
			handleError(ctx, err)
			return
		}
		handleOK(ctx, res.Expose())
	})
}
