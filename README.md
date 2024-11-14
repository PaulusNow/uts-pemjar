# Real-Time Donation System

This project involves a real-time donation system utilizing three protocols: TCP, UDP, and WebSocket. Below are the steps to run the application, depending on which protocol you want to use.

## WebSocket

WebSocket is implemented separately from TCP and UDP. To run the WebSocket server and the Next.js frontend, follow these steps:

1. **Switch to the `master` branch** to use the WebSocket functionality.
2. **Download the WebSocket files** from the `master` branch.
3. **For the backend**: Download the backend code from the `main` branch and combine the files into the same directory.
4. **Install dependencies for the frontend**: 
   First, you need to install Next.js in the frontend directory. Open a terminal in your project folder and run:
   
   ```bash
   npm install next
5. **Run the WebSocket server**:
   Open two terminals:
   
   - **Terminal 1**: Run the Go backend server:
     ```bash
     go run main.go
     ```

   - **Terminal 2**: Start the Next.js frontend:
     ```bash
     npm run dev
     ```

6. **Access the application**: Open [http://localhost:3000](http://localhost:3000) in your browser to see the real-time donation system in action.

## TCP and UDP

For the TCP and UDP functionality, follow the steps below:

1. **Switch to the `tcp-&-udp` branch** for the TCP and UDP features.
2. **Download the TCP and UDP files** from the `tcp-&-udp` branch.
3. **Run the application in three terminals**:
   
   - **Terminal 1**: Run the Go server (`server.go`):
     ```bash
     go run server.go
     ```

   - **Terminal 2**: Run the Go client (`client.go`):
     ```bash
     go run client.go
     ```

   - **Terminal 3**: Run the Go top-up functionality (`topup.go`):
     ```bash
     go run topup.go
     ```

This will allow you to interact with the system using TCP for registration and login, UDP for top-up functionality, and a simple terminal interface for managing donations.

## Additional Information

### Learn More

- [Next.js Documentation](https://nextjs.org/docs) - Learn more about Next.js features and APIs.
- [Learn Next.js](https://nextjs.org/learn) - Interactive tutorial for Next.js.
- [Vercel Deployment Documentation](https://nextjs.org/docs/app/building-your-application/deploying) - How to deploy your Next.js app on Vercel.

### Deploy on Vercel

The easiest way to deploy your Next.js app is through the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme).

### Video Tutorial

For a detailed walkthrough on setting up and running the donation system, watch the video on YouTube:

[Watch the video here](https://youtu.be/JEMGIW0JDzo)
