package io;

import io.xjar.XCryptos;

public class ByzkCaMain {

    public static void main(String[] args) throws Exception {

        String WORK_DIR = System.getProperty("user.dir");
        String LIB_DIR = WORK_DIR + "/lib/";
        String camanage = LIB_DIR + "test.jar";
        String camanageEncPath = LIB_DIR + "test-enc.jar";

        XCryptos.encryption()
                .from(camanage)
                .use("sjfadfare3adfadf")
                .include("/**")
                .exclude("/static/**/*")
                .exclude("/conf/*")
                .to(camanageEncPath);
    }
}
