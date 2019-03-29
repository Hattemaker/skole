/**
 * Superclass.
 */
public class Menneske {

    protected String navn;
    protected int liv = 100;
    protected int styrke = 1;

    /**
     * Constructor for objekter av klassen Menneske.
     * @param navn Navnet på mennesket.
     */
    public Menneske(String navn) {
        this.navn = navn;

    }

    /**
     * Skader motstander med mengde styrke den som skader har.
     * @param motstander Mostanderen til den som gjør skade.
     */
    public void skad(Menneske motstander) {
        motstander.reduserLiv(styrke);
        System.out.println(motstander.getNavn() + " ble skadet ");

    }

    /**
     * Reduserer antall liv med antall styrke motstander har.
     * @param styrke Styrken den som reduserer liv har.
     */
    public void reduserLiv(int styrke) {
       liv -= styrke;

    }

    /**
     * Printer resterende liv.
     */
    public void antallLiv(){
        System.out.println("Resterende liv er: " + liv + "%");
    }

    /**
     *
     * @return Navnet på objektet.
     */
    public String getNavn(){
        return navn;
    }


}