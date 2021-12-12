import 'dart:io';

/*TO WRITE*/
String strStop = "docker-compose -f dc.d down";
String strDockerCompose =
    "version: '3.5'\nservices:\n    hapi-fhir-server:\n      image: hopsiia/fhir-isih:latest\n      container_name: hapi-fhir-server\n      hostname: hapi-fhir-server\n      restart: on-failure\n      ports:\n        - '8181:8080'\n    hapi-fhir-pg:\n      image: postgres:12\n      container_name: hapi-fhir-pg\n      hostname: hapi-db\n      restart: always\n      environment:\n        - POSTGRES_DB=hapi_r4\n        - POSTGRES_USER=postgres\n        - POSTGRES_PASSWORD=admin\n      command: postgres -c 'max_connections=200'\n      ports:\n        - '5432:5432'\n      volumes:\n       - './pgdatafhir:/var/lib/postgresql/data'\nnetworks:\n    hapi-bridge:\n      name: hapi-network";

/**********************************
 *      MAIN
 **********************************/
void main(List<String> arguments) {
  message();
  writeOnDisk(strStop, "s.d");
  writeOnDisk(strDockerCompose, "dc.d");
  eraseOnDisk("r.d");
  running("/bin/sh", "./s.d");
  eraseOnDisk("s.d");
  eraseOnDisk("dc.d");
}

/**********************************
 * Ecriture sur disque du fichier...
 **********************************/
void writeOnDisk(String valeurs, String nomfichier) {
  try {
    var file = File(nomfichier)
        .writeAsStringSync(valeurs, mode: FileMode.write, flush: true);
  } on FileSystemException {
    //do nothing
    print("Write Error...");
  } on IOException {
    //DO Nothing
    print("File Access Error...");
  }
}

/***********************************
 * Supression sur disque...
 ***********************************/
void eraseOnDisk(String nomfichier) {
  try {
    var file = File(nomfichier).deleteSync();
  } on FileSystemException {
    //do nothing
    print("Suppress Error...");
  } on IOException {
    //DO Nothing
    print("File Access Error...");
  }
}

/**************************************
 *  Lancement de l'application...
 **************************************/
Future<void> running(String cmd, String opt) async {
  ProcessResult result;
  try {
    result = await Process.run(cmd, [opt]);
    print(result.stdout.toString());
  } on Exception {
    print("Execution impossible...");
  }
}

/*************************************
 * Message accueil...
 *************************************/
void message() {
  String strMessage = """
  ==============================
 Arret du serveur FHIR-ISIH
==============================""";
  print(strMessage);
}
