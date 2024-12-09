package org.valhalla.user.api.route

import io.ktor.server.application.*
import org.akrck02.valhalla.core.dal.configuration.getCurrentDatabaseConfiguration
import org.akrck02.valhalla.core.dal.database.Databases
import org.akrck02.valhalla.core.dal.database.Mongo
import org.akrck02.valhalla.core.dal.service.user.UserDataAccess

fun Application.configureRouting() {

    val mongo = Mongo().also { it.connect(getCurrentDatabaseConfiguration()) }
    val database = mongo.getDatabase(Databases.Valhalla)
    val userDataAccess = UserDataAccess(database = database)

    configureUserRoutes(userDataAccess)
}


