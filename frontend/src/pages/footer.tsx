
import exachan from '../img/exachan.png';
import pravdorubs from '../img/pravdorubs.png';

export default function Footer() {
    return (
        <footer id="footer">
            <a href="http://exachan.com/">
                <img src={exachan} alt="Эксачан" title="Эксачан" />
            </a>
            <a href="http://pravdorubs.ru/">
                <img src={pravdorubs} alt="Правдорубы" title="Правдорубы" />
            </a>
            <a href="http://exachan.com/exaba/">эксаба 1.0.0 for Go</a>
        </footer>
    );
}