///usr/bin/env jbang "$0" "$@" ; exit $?
//DEPS jakarta.json.bind:jakarta.json.bind-api:3.0.1
//DEPS org.eclipse:yasson:3.0.3
//DEPS jakarta.json:jakarta.json-api:2.1.3
//DEPS commons-io:commons-io:2.16.1

// mvn dependenncy:get -DgroupId=jakarta.json.bind -DartifactId=jakarta.json.bind-api -Dversion=3.0.1
// mvn dependency:get -DgroupId=org.eclipse -DartifactId=yasson -Dversion=3.0.3
// mvn dependency:get -DgroupId=jakarta.json -DartifactId=jakarta.json-api -Dversion=2.1.3

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.SocketException;
import java.net.URL;

import java.util.List;
import java.util.Map;
import java.util.Base64;

import java.time.LocalDateTime;

import java.io.FileWriter;
import java.io.IOException;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.Path;

import java.util.regex.Matcher;
import java.util.regex.Pattern;


import java.io.FileOutputStream;
import java.io.InputStream;

import jakarta.json.bind.Jsonb;
import jakarta.json.bind.JsonbBuilder;
import jakarta.json.bind.JsonbBuilder;

import org.apache.commons.io.FileUtils;

public class SApiOpm {

    private static String opmURL = "https://gist.githubusercontent.com/funkyhippo/1d40bd5dae11e03a6af20e5a9a030d81/raw/?";

    public static void main(String... args) {
        try {
            var response = getRequest();
            if (response == null) {
                System.out.println("Something went wrong with response, exiting!");
                System.exit(1);
            }

            Jsonb jsonb = JsonbBuilder.create();
            Opm opm = jsonb.fromJson(response, Opm.class);

            printInfo(opm);
           getPics(opm.chapters);
        }
        catch(jakarta.json.bind.JsonbException je) {
            System.out.println(je);
            je.printStackTrace();
        }
        catch (Exception e) {
            System.out.println(e);
            e.printStackTrace();
        }

    }

    private static void getPics(Map<String, Chapter> chapters) {
        var home = System.getProperty("user.home");
        chapters.forEach((k, v) -> {
                var chapterDir = Paths
                    .get(home, "Downloads", "opm", k)
                    .toString()
                    .replace(".", "_");

                System.out.println();
                System.out.println("Chapter: " + k);
                System.out.printf("%-2sDirectory: %s\n", "", chapterDir);

                v.groups.get("/r/OnePunchMan").forEach(link -> {
                        try {
                            Thread.sleep(3000);
                            var finalLink = getDecodedLink(link);
                            var picName = new String(URI.create(finalLink).getPath()).substring(1);
                            var filepath = Paths.get(chapterDir, picName);

                            if(!Files.exists(filepath)) {
                                System.out.printf("%-4surl: %s - filepath: %s\n", "", finalLink, filepath);
                                try {
                                    Files.createDirectories(Paths.get(chapterDir));

                                    var finalLinkUrl = new URL(finalLink).openConnection();
                                    finalLinkUrl.setRequestProperty("User-Agent", "Mozilla/5.0");
                                    FileUtils.copyInputStreamToFile(finalLinkUrl.getInputStream(),
                                                                    new File(filepath.toString()));
                                } catch(SocketException se) {
                                    System.out.println("Error: " + se.getMessage());
                                    se.printStackTrace();
                                }
                                catch(IOException ioe) {
                                    System.out.println("Error: error copying file: " + ioe.getMessage());
                                    ioe.printStackTrace();
                                }
                            }
                        } catch(InterruptedException e){
                            System.out.println(e.getMessage());
                            e.printStackTrace();
                        }
                    });
            });
    }

    private static String chapterZeroAppend(String meh) {
        String regex = "\\b(\\d)_";
        Pattern pattern = Pattern.compile(regex);
        Matcher matcher = pattern.matcher(meh);
        if(matcher.find()) {
            return "0" + meh;
        }

        return meh;
    }

    private static void printInfo(Opm opm) {
        System.out.printf("\ntitle: %s\n", opm.title);
        System.out.printf("latest cover: %s\n", getDecodedLink(opm.cover));
    }

	private static String getRequest() {
      var jsonLocation = "/tmp/opm"+ ".json"; // + LocalDateTime.now()
      var file = new File(jsonLocation);
      if(file.exists()) {
          System.out.println("WARNING: JSON file already downloaded, imma using it!");
          try {
              return String.join(System.lineSeparator(), Files.readAllLines(Paths.get(jsonLocation)));
          } catch(IOException ioe) {
              System.out.println("An error occurred while saving the JSON data: " + ioe.getMessage());
              ioe.printStackTrace();
}
      }

      HttpResponse<String> response = null;
		var client = HttpClient.newHttpClient();
      var request = HttpRequest
            .newBuilder()
            .uri(URI.create(opmURL))
            .build();

        try {
            response = client.send(request, HttpResponse.BodyHandlers.ofString());
        } catch(InterruptedException ie) {
            System.out.println("An error occurred while saving the JSON data: " + ie.getMessage());
            ie.printStackTrace();
        } catch (IOException e) {
            System.out.println("An error occurred while saving the JSON data: " + e.getMessage());
            e.printStackTrace();
        }

        try {
            var filew = new FileWriter(jsonLocation);
            filew.write(response.body());
            filew.close();
        } catch(IOException ioe) {
            System.out.println("An error occurred while saving the JSON data: " + ioe.getMessage());
            ioe.printStackTrace();
        }

        return response.body();
	}

    private static String getDecodedLink(String link) {
        String result = null;
        var uLink = URI.create(link).getPath().replaceAll("/v1/image/", "");
        result = new String(Base64.getDecoder().decode(uLink));

        return result;
    }

    public record Chapter(
                          Map<String, List<String>> groups,
                          long lastUpdated,
                          String title,
                          String volume) {
    }

    public record Opm(
                      String cover,
                      String description,
                      String title,
                      Map<String, Chapter> chapters) {
    }

}
