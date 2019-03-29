/**
 * Kriger extends Menneske.
 */
public class Kriger extends Menneske {

    private int ekstraStyrke = 3; // Krigere sin ekstra styrke.

    /**
     * Consructor for klassen Kriger.
     * @param navn
     */
    public Kriger(String navn) {
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
