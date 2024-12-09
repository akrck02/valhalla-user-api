package org.valhalla.user.api.route

import io.ktor.http.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.akrck02.valhalla.core.sdk.error.ErrorCode
import org.akrck02.valhalla.core.sdk.model.exception.ServiceException

/**
 * Run code returning exceptions in JSON format
 * @param call The routing call
 * @param code The code to execute
 */
suspend fun runSecure(
    call: RoutingCall,
    code: suspend () -> Unit
) {
    try {
        code()
    } catch (e: ServiceException) {
        call.respond(
            status = HttpStatusCode(value = e.status.code, description = e.message),
            message = Json.encodeToString(e)
        )
    } catch (e: Exception) {
        call.respond(
            status = HttpStatusCode.InternalServerError,
            message = Json.encodeToString(
                ServiceException(
                    status = org.akrck02.valhalla.core.sdk.model.http.HttpStatusCode.InternalServerError,
                    code = ErrorCode.UnexpectedError,
                    message = e.message.orEmpty()
                )
            )
        )
    }
}
