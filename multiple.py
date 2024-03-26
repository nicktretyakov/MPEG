import cv2

# Initialize the video capture objects for each camera
cap0 = cv2.VideoCapture(0)  # Default camera
cap1 = cv2.VideoCapture(1)  # Second camera

while True:
    # Read frames from each camera
    ret0, frame0 = cap0.read()
    ret1, frame1 = cap1.read()

    # Display frames in separate windows
    cv2.imshow('Camera 0', frame0)
    cv2.imshow('Camera 1', frame1)

    # Break the loop if 'q' is pressed
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# Release the video capture objects and close all windows
cap0.release()
cap1.release()
cv2.destroyAllWindows()