import {Link} from "react-router-dom";
import {Page} from "../../../components";

export default class Home extends Page {
  render() {
    return (
      <div>
        <h1>Home</h1>

        <ul>
          <li>
            <Link to="/projects">Projects - Home</Link>
            <Link to="/security">Security - Home</Link>
            <Link to={"/security/login"}>Security - Login</Link>
          </li>
        </ul>
      </div>
    );
  }
}
