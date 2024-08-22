package io.xjar.boot;

import io.xjar.XLauncher;
import io.xjar.jar.XJarClassLoader;
import org.springframework.boot.loader.launch.Archive;
import org.springframework.boot.loader.launch.ExecutableArchiveLauncher;
import org.springframework.boot.loader.launch.JarLauncher;

import java.io.File;
import java.lang.reflect.Method;
import java.net.URI;
import java.net.URL;
import java.net.URLClassLoader;
import java.security.CodeSource;
import java.security.ProtectionDomain;
import java.util.Collection;
import java.util.jar.Attributes;
import java.util.jar.JarFile;
import java.util.jar.Manifest;

/**
 * Spring-Boot Jar 启动器
 *
 * @author Payne 646742615@qq.com
 * 2018/11/23 23:06
 */
//public class XJarLauncher extends ExecutableArchiveLauncher {
public class XJarLauncher extends JarLauncher {
    private final XLauncher xLauncher;

    public XJarLauncher(String... args) throws Exception {
        this.xLauncher = new XLauncher(args);
    }

    public static void main(String[] args) throws Exception {
        new XJarLauncher(args).launch();
    }

    public void launch() throws Exception {
        launch(xLauncher.args);
    }

    @Override
    protected ClassLoader createClassLoader(Collection<URL> urls) throws Exception {
        URL[] urlArray = urls.toArray(new URL[0]);
        return new XBootClassLoader(urlArray, this.getClass().getClassLoader(), xLauncher.xDecryptor, xLauncher.xEncryptor, xLauncher.xKey);
    }


    @Override
    protected boolean isIncludedOnClassPath(Archive.Entry entry) {
        return isLibraryFileOrClassesDirectory(entry);
    }

    static boolean isLibraryFileOrClassesDirectory(Archive.Entry entry) {
        String name = entry.name();
        return entry.isDirectory() ? name.equals("BOOT-INF/classes/") : name.startsWith("BOOT-INF/lib/");
    }

    @Override
    protected String getEntryPathPrefix() {
        return "BOOT-INF/";
    }
}
