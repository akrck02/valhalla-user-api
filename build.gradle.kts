import java.io.FileInputStream
import java.util.*

// Variables section
val localProperties: Properties = Properties().apply {
    load(FileInputStream(File(rootProject.rootDir, "local.properties")))
}

group = providers.gradleProperty("organization").getOrElse("")
version = providers.gradleProperty("valhalla.api.user.version").getOrElse("")

val ktorVersion: String = providers.gradleProperty("kot.version").getOrElse("")
val jdkVersion: Int = providers.gradleProperty("jdk.version").getOrElse("").toInt()
val mavenServerName: String = localProperties.getProperty("maven.server.name")
val mavenServerUrl: String = localProperties.getProperty("maven.server.url")
val mavenServerUser: String = localProperties.getProperty("maven.server.user")
val mavenServerPassword: String = localProperties.getProperty("maven.server.password")
val valhallaSdkVersion: String = providers.gradleProperty("valhalla.core.sdk.version").getOrElse("")
val valhallaDalVersion: String = providers.gradleProperty("valhalla.core.dal.version").getOrElse("")

// Plugins section
plugins {
    alias(libs.plugins.kotlin.jvm)
    alias(libs.plugins.ktor)
    id("org.jetbrains.kotlin.plugin.serialization") version "2.1.0"
}

// Ktor Application section
application {
    mainClass.set("io.ktor.server.netty.EngineMain")

    val isDevelopment: Boolean = project.ext.has("development")
    applicationDefaultJvmArgs = listOf("-Dio.ktor.development=$isDevelopment")
}

// Repositories
val selfHostedRepository = repositories.maven {
    name = mavenServerName
    url = uri(mavenServerUrl)
    credentials {
        username = mavenServerUser
        password = mavenServerPassword
    }
    isAllowInsecureProtocol = true
}
repositories {
    mavenCentral()
    selfHostedRepository
    mavenLocal()
}

// Compilation section
kotlin {
    jvmToolchain(jdkVersion)
}

java {
    withSourcesJar()
    withJavadocJar()
}

// Publishing section
dependencies {

    // Ktor dependencies
    implementation(libs.ktor.server.core)
    implementation(libs.ktor.server.swagger)
    implementation(libs.ktor.server.cors)
    implementation(libs.ktor.server.netty)
    implementation(libs.logback.classic)
    implementation(libs.ktor.server.config.yaml)
    implementation("io.ktor:ktor-server-resources:$ktorVersion")

    // Kotlin coroutine dependency
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.6.4")

    // MongoDB Kotlin driver dependency
    implementation("org.mongodb:mongodb-driver-kotlin-coroutine:4.10.1")

    // Serialization
    implementation("io.ktor:ktor-server-content-negotiation:$ktorVersion")
    implementation("io.ktor:ktor-client-content-negotiation:$ktorVersion")
    implementation("io.ktor:ktor-serialization-kotlinx-json:$ktorVersion")
    implementation("io.ktor:ktor-serialization-kotlinx-protobuf:$ktorVersion")

    // Valhalla dependencies
    implementation("org.akrck02:valhalla.core.sdk:$valhallaSdkVersion")
    implementation("org.akrck02:valhalla.core.dal:$valhallaDalVersion")

    // Testing
    testImplementation(libs.ktor.server.test.host)
    testImplementation(libs.kotlin.test.junit)

}
