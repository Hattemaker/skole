/**
 * Bueskytter extends Menneske.
 */
public class Bueskytter extends Menneske {

    private int ekstraStyrke = 8; // Bueskyttere sin ekstra skade.

    /**
     * constructor for klassen Bueskytter.
     * @param navn
     */
    public Bueskytter(String navn) {
        super(navn);
    }


    /**
     * Metode som overskriver "skad"-metoden i "Menneske"-klassen, og dermed ggjør det mulig å skade mer.
     * @param motstander
     */
    @Override
    public void skad (Menneske motstander){
        int totalStyrke = styrke + ekstraStyrke;
        motstander.reduserLiv(styrke + ekstraStyrke);
        System.out.println(this.navn +" skadet " + motstander.getNavn() + " med en styrke på " + totalStyrke + ".");

    }
}