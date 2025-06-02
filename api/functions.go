package api

type genericFunctionBody struct {
	Function string `json:"function"`
}

const HealthCheckFunction = "HealthCheck"

const VerifyAuthTokenFunction = "VerifyAuthenticationToken"

const PasswordlessLoginFunction = "PasswordlessLogin"

const PasswordLoginFunction = "PasswordLogin"

const QueryServerStateFunction = "QueryServerState"

const GetServerOptionsFunction = "GetServerOptions"

const GetAdvancedGameSettingsFunction = "GetAdvancedGameSettings"

const ApplyAdvancedGameSettingsFunction = "ApplyAdvancedGameSettings"

const ClaimServerFunction = "ClaimServer"

const RenameServerFunction = "RenameServer"

const SetClientPasswordFunction = "SetClientPassword"

const SetAdminPasswordFunction = "SetAdminPassword"

const SetAutoLoadSessionNameFunction = "SetAutoLoadSessionName"

const RunCommandFunction = "RunCommand"

const ShutdownFunction = "Shutdown"

const ApplyServerOptionsFunction = "ApplyServerOptions"

const CreateNewGameFunction = "CreateNewGame"

const SaveGameFunction = "SaveGame"

const DeleteSaveFileFunction = "DeleteSaveFile"

const DeleteSaveSessionFunction = "DeleteSaveSession"

const EnumerateSessionsFunction = "EnumerateSessions"

const LoadGameFunction = "LoadGame"

const UploadSaveGameFunction = "UploadSaveGame"

const DownloadSaveGameFunction = "DownloadSaveGame"
