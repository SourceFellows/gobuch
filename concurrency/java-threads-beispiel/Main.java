import java.util.concurrent.Semaphore;
import java.util.concurrent.atomic.AtomicBoolean;

public class Main {

	static int count = 0;

	public static void main(String[] args) {

		new Thread(new Runnable() {
			@Override
			public void run() {
				while (true) {
					count++;
					try {
						new Thread(new Runnable() {
							@Override
							public void run() {
								while (true) {
									try {
										Thread.sleep(120000);
									} catch (InterruptedException e) {
										e.printStackTrace();
									}
									System.out.println("Thread Hello " + count);
								}
							}
						}).start();
					} catch (Throwable e) {
						e.printStackTrace();
						System.exit(1);
					}
					System.out.println("create native thread "+ count);
				}
			}
		}).start();

	}

}
