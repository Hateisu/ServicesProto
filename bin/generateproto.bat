@echo off
setlocal enabledelayedexpansion

REM Папки

set ROOT_DIR=D:\Programming\Projects\Auck\backend

REM Сервисы
set SERVICE[0]=auth
set SERVICE[1]=marketplace
set SERVICE[2]=gmail
set SERVICE[3]=kafkamessages
set SERVICE[4]=filemanager
set SERVICE[5]=redis_messages

REM Суфиксы
set GEN_SUFIX=%ROOT_DIR%\proto\protos_gen
set INCLUDE_DIR=%ROOT_DIR%\proto\include
set SERVICES_DIR=%ROOT_DIR%\services

set AUTH_PROTO_DIR=D:\Programming\Projects\Auck\backend\proto\auth
set AUTH_OUT_DIR=D:\Programming\Projects\Auck\backend\proto\protos_gen\auck.authservice1.v1

set MARKETPLACE_PROTO_DIR=D:\Programming\Projects\Auck\backend\proto\marketplace
set MARKETPLACE_OUT_DIR=D:\Programming\Projects\Auck\backend\proto\protos_gen\auck.marketplaceservice1.v1

set MAILING_PROTO_DIR=D:\Programming\Projects\Auck\backend\proto\gmailservice
set MAILING_OUT_DIR=D:\Programming\Projects\Auck\backend\proto\protos_gen\auck.gmailservice1.v1

set INCLUDE_DIR=D:\Programming\Projects\Auck\backend\proto\include
set OUT_DIR=D:\Programming\Projects\Auck\backend\proto\protos_gen

REM Генерация
REM protoc ^
REM  -I=%AUTH_PROTO_DIR% ^
REM  -I=%INCLUDE_DIR% ^
REM  --go_out=%AUTH_OUT_DIR% --go_opt=paths=source_relative ^
REM  --go-grpc_out=%AUTH_OUT_DIR% --go-grpc_opt=paths=source_relative ^
REM  --grpc-gateway_out=%AUTH_OUT_DIR% --grpc-gateway_opt=paths=source_relative ^
REM  --openapiv2_out=%OUT_DIR% ^
REM  %AUTH_PROTO_DIR%\*.proto
REM protoc ^
REM  -I=%MARKETPLACE_PROTO_DIR% ^
REM  -I=%INCLUDE_DIR% ^
REM  --go_out=%MARKETPLACE_OUT_DIR% --go_opt=paths=source_relative ^
REM  --go-grpc_out=%MARKETPLACE_OUT_DIR% --go-grpc_opt=paths=source_relative ^
REM  --grpc-gateway_out=%MARKETPLACE_OUT_DIR% --grpc-gateway_opt=paths=source_relative ^
REM  --openapiv2_out=%OUT_DIR% ^
REM  %MARKETPLACE_PROTO_DIR%\*.proto
REM 
REM protoc ^
REM  -I=%MAILING_PROTO_DIR% ^
REM  -I=%INCLUDE_DIR% ^
REM  --go_out=%MAILING_OUT_DIR% --go_opt=paths=source_relative ^
REM  --go-grpc_out=%MAILING_OUT_DIR% --go-grpc_opt=paths=source_relative ^
REM  --grpc-gateway_out=%MAILING_OUT_DIR% --grpc-gateway_opt=paths=source_relative ^
REM  --openapiv2_out=%OUT_DIR% ^
REM  %MAILING_PROTO_DIR%\*.proto

REM Генерация 
set COPY_TARGETS=AuthService MarketplaceService GmailService APIGateway FileManagerService
for %%i in (0,1,2,3,4,5) do (
    set NAME=!SERVICE[%%i]!
    set INPUT_DIR=%ROOT_DIR%\proto\!NAME!
    if "!NAME!" == "kafkamessages"  (
    set OUT_DIR=%GEN_SUFIX%\auck.!NAME!.v1
    ) else if /i "!NAME!" == "redis_messages" (
    set OUT_DIR=%GEN_SUFIX%\auck.!NAME!.v1
    ) else (
    set OUT_DIR=%GEN_SUFIX%\auck.!NAME!service1.v1
    )

    rem set DEST_DIR=%SERVICES_DIR%\!NAME!\internal\proto\auck.!NAME!service1.v1

    echo.
    echo Generating proto for: !NAME! from !INPUT_DIR! to !OUT_DIR!
    protoc ^
     -I=!INPUT_DIR! ^
     -I=%INCLUDE_DIR% ^
     --go_out=!OUT_DIR! --go_opt=paths=source_relative ^
     --go-grpc_out=!OUT_DIR! --go-grpc_opt=paths=source_relative ^
     --grpc-gateway_out=!OUT_DIR! --grpc-gateway_opt=paths=source_relative ^
     --openapiv2_out=%GEN_SUFIX% ^
     !INPUT_DIR!\*.proto
)
    for %%T in (%COPY_TARGETS%) do (
        set DEST_DIR=%SERVICES_DIR%\%%T\internal\proto
        if exist "!DEST_DIR!" (
            echo Copying "!GEN_SUFIX!" to "!DEST_DIR!"
            robocopy "!GEN_SUFIX!" "!DEST_DIR!" /E /NFL /NDL /NJH /NJS 
        ) else (
            echo Dir not found: !DEST_DIR!
        )
    )
    set SWAGGER_DEST_DIR=%SERVICES_DIR%\APIGateway\swagger
        if exist "!SWAGGER_DEST_DIR!" (
            echo Copying "!GEN_SUFIX!" to "!SWAGGER_DEST_DIR!"
            robocopy "!GEN_SUFIX!" "!SWAGGER_DEST_DIR!" /E /NFL /NDL /NJH /NJS 
        ) else (
            echo Dir not found: !SWAGGER_DEST_DIR!
        )

echo.
echo Done