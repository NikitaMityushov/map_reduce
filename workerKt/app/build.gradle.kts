import com.google.protobuf.gradle.id

plugins {
    `java-library`
    kotlin("jvm")
    id("com.google.protobuf") version "0.9.4"
}

kotlin {
    jvmToolchain(8)
}

repositories {
    google()
    gradlePluginPortal()
    mavenCentral()
}

protobuf {
    protoc {
        artifact = "com.google.protobuf:protoc:4.27.3"
    }

    plugins {
        id("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:1.62.2"
        }
        id("grpckt") {
            artifact = "io.grpc:protoc-gen-grpc-kotlin:1.4.1:jdk8@jar"
        }
    }

    generateProtoTasks {
        all().forEach { task ->
            task.plugins {
                id("grpc")
                id("grpckt")
            }

            task.builtins {
                id("kotlin")
            }
        }
    }
}

dependencies {
    protobuf(project(":protos"))
    implementation("com.google.protobuf:protobuf-kotlin:4.27.3")
    api("io.grpc:grpc-protobuf:1.62.2")
    api("com.google.protobuf:protobuf-java-util:4.27.3")
    api("com.google.protobuf:protobuf-kotlin:4.27.3")
    api("io.grpc:grpc-kotlin-stub:1.4.1")
    api("io.grpc:grpc-stub:1.62.2")

    testImplementation(kotlin("test"))
}
