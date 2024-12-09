package org.valhalla.user.api.route

import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.resources.*
import io.ktor.server.resources.patch
import io.ktor.server.resources.post
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.akrck02.valhalla.core.dal.configuration.getCurrentDatabaseConfiguration
import org.akrck02.valhalla.core.dal.database.Databases
import org.akrck02.valhalla.core.dal.database.Mongo
import org.akrck02.valhalla.core.dal.service.user.UserDataAccess
import org.akrck02.valhalla.core.sdk.model.exception.ServiceException
import org.akrck02.valhalla.core.sdk.model.user.User

fun Application.configureRouting() {

    val mongo = Mongo().also { it.connect(getCurrentDatabaseConfiguration()) }
    val database = mongo.getDatabase(Databases.Valhalla)
    val userDataAccess = UserDataAccess(database = database)

    routing {

        get<UserRoute.Id> {
            runSecure(call) {
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.get(it.id)
                )
            }
        }

        post<UserRoute.Register> {
            runSecure(call) {

                val body = call.receive<User>()
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.register(body)
                )
            }
        }

        patch<UserRoute.Id.Update> {
            runSecure(call) {
                val body = call.receive<User>()
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.update(it.parent.id, body)
                )
            }
        }

    }
}


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
    }
}
