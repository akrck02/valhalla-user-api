package org.valhalla.user.api.route

import io.ktor.resources.*

@Resource("/users")
class UserRoute {

    @Resource("register")
    class Register(val parent: UserRoute)

    @Resource("{id}")
    class Id(val parent: UserRoute, val id: String?) {

        @Resource("edit")
        class Update(val parent: Id)

    }
}