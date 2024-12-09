package org.valhalla.user.api.route

import io.ktor.http.*
import io.ktor.resources.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.resources.*
import io.ktor.server.resources.patch
import io.ktor.server.resources.post
import io.ktor.server.response.*
import io.ktor.server.routing.*
import org.akrck02.valhalla.core.dal.service.user.UserDataAccess
import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * User route typed tree
 */
@Suppress("unused")
@Resource("/users")
class UserRoute {

    @Resource("register")
    class Register(val parent: UserRoute)

    @Resource("login")
    class Login(val parent: UserRoute, val user: User?) {

        @Resource("auth")
        class Auth(val parent: Login, val auth: String?)

    }

    @Resource("{id}")
    class Id(val parent: UserRoute, val id: String?)

    @Resource("emails/{email}")
    class Email(val parent: UserRoute, val email: String?)

    @Resource("validate")
    class Validate(val parent: UserRoute, val code: String?)

}

/**
 * Register user routes
 */
fun Application.configureUserRoutes(userDataAccess: UserDataAccess) {
    routing {
        post<UserRoute.Register> {
            runSecure(call) {

                val body = call.receive<User>()
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.register(body)
                )
            }
        }

        get<UserRoute.Id> {
            runSecure(call) {
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.get(it.id)
                )
            }
        }

        patch<UserRoute.Id> {
            runSecure(call) {
                val body = call.receive<User>()
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.update(it.id, body)
                )
            }
        }

        delete<UserRoute.Id> {
            runSecure(call) {
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.delete(it.id)
                )
            }
        }

        get<UserRoute.Email> {
            runSecure(call) {
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.getByEmail(it.email)
                )
            }
        }

        get<UserRoute.Validate> {
            runSecure(call) {
                call.respond(
                    status = HttpStatusCode.OK,
                    message = userDataAccess.validateAccount(it.code)
                )
            }
        }

    }
}