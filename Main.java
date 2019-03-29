public class Main {

    public static void main(String[] args) {
        Menneske kriger1 = new Kriger("Frank");
        Menneske bueskytter1 = new Bueskytter("Lisa");

/**
 * Test-interaksjon.
 */
        kriger1.skad(bueskytter1);
        bueskytter1.antallLiv();
        System.out.println();

        bueskytter1.skad(kriger1);
        kriger1.antallLiv();
        System.out.println();

    }
}
