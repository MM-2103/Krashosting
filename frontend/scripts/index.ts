import axios from "axios";

interface Session {
  token: string;
  user: User;
}
interface User {
  success: boolean;
  session: Session;
}

const login = async (username: string, password: string): Promise<boolean> => {
  const response = await axios.post<User>("http://localhost:3000/login", {
    username,
    password,
  });
  return response.data.success;
};
