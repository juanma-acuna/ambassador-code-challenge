import { render, screen } from "@testing-library/react";
import App from "./App";

test("render the main page", () => {
  render(<App />);
  const linkElement = screen.getByText(/Services Viewer/i);
  expect(linkElement).toBeInTheDocument();
});
