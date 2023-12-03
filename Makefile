CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.2
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
MOCKGEN=${BINDIR}/mockgen_${GOVER}
PROTOCBIN=${BINDIR}/protoc_${GOVER}

# ==============================================================================
# Tools commands

install-mockgen: bindir
	test -f ${MOCKGEN} || \
		(GOBIN=${BINDIR} go install github.com/golang/mock/mockgen@v1.6.0 && \
		mv ${BINDIR}/mockgen ${MOCKGEN})
.PHONY: install-mockgen

gen-mocks: install-mockgen
	rm -r order/internal/mocks/repo/
	${MOCKGEN} -source=order/internal/infrastructure/interfaces/order.go -destination=order/internal/mocks/repo/order_mocks.go
	${MOCKGEN} -source=user/internal/infrastructure/interfaces/user.go -destination=user/internal/mocks/repo/user_mocks.go
	${MOCKGEN} -source=product/internal/infrastructure/interfaces/product.go -destination=product/internal/mocks/repo/product_mocks.go
	${MOCKGEN} -source=product/internal/infrastructure/interfaces/discount.go -destination=product/internal/mocks/repo/discount_mocks.go

	${MOCKGEN} -source=user/internal/usecase/user.go -destination=user/internal/mocks/usecase/user_mocks.go
	${MOCKGEN} -source=product/internal/usecase/product.go -destination=product/internal/mocks/usecase/product_mocks.go
.PHONY: gen-mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})
.PHONY: install-lint

lint: install-lint
	${LINTBIN} run
.PHONY: install-lint

precommit: format lint
	echo "OK"
.PHONY: precommit

bindir:
	mkdir -p ${BINDIR}
.PHONY: bindir

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})
.PHONY: install-smartimports

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks
.PHONY: format

install-protoc: bindir
	test -f ${PROTOCBIN} || \
		(GOBIN=${BINDIR} go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
		mv ${BINDIR}/protoc-gen-go-grpc ${PROTOCBIN})
.PHONY: install-protoc

protoc-order:
	mkdir -p proto/gen/order
	protoc -I proto \
		--go_out proto/gen/order \
		--go_opt paths=source_relative \
		--go-grpc_out proto/gen/order \
		--go-grpc_opt paths=source_relative \
		proto/order.proto
.PHONY: protoc-order

protoc-user:
	mkdir -p proto/gen/user
	protoc -I proto \
		--go_out proto/gen/user \
		--go_opt paths=source_relative \
		--go-grpc_out proto/gen/user \
		--go-grpc_opt paths=source_relative \
		proto/user.proto
.PHONY: protoc-user

protoc-cart:
	mkdir -p proto/gen/cart
	protoc -I proto \
		--go_out proto/gen/cart \
		--go_opt paths=source_relative \
		--go-grpc_out proto/gen/cart \
		--go-grpc_opt paths=source_relative \
		proto/cart.proto
.PHONY: protoc-cart

protoc-product:
	mkdir -p proto/gen/product
	protoc -I proto \
		--go_out proto/gen/product \
		--go_opt paths=source_relative \
		--go-grpc_out proto/gen/product \
		--go-grpc_opt paths=source_relative \
		proto/product.proto
.PHONY: protoc-product

protoc-gateway:
	mkdir -p proto/gen/gateway
	protoc -I proto \
		--go_out proto/gen/gateway \
		--go_opt paths=source_relative \
		--go-grpc_out proto/gen/gateway \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out proto/gen/gateway \
    	--grpc-gateway_opt paths=source_relative \
		--openapiv2_out docs \
		--openapiv2_out gateway/internal/app \
		proto/gateway.proto
.PHONY: protoc-gateway

protoc-all: protoc-order protoc-gateway protoc-user protoc-cart protoc-product
	@echo All Protobufs Generated
.PHONY: protoc-all

# ==============================================================================
# Tests commands

test:
	go test -v -race -count=1 ./.../internal/...
.PHONY: test

test-100:
	go test -v -race -count=100 ./.../internal/...
.PHONY: test

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./.../internal/...
	go tool cover -html=coverage.out -o coverage.html
	xdg-open coverage.html
	rm coverage.out
	rm coverage.html
.PHONY: cover
