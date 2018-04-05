using System.Collections;
using System.Collections.Generic;
using UnityEngine;
//Camera Controls: Zoom, Accelerated Rotation, Re-focus
public class CameraController : MonoBehaviour {

    public GameObject freeLookCam;
    private Cinemachine.CinemachineFreeLook freeLook;

	void Start () {
        freeLook = freeLookCam.GetComponent<Cinemachine.CinemachineFreeLook>();
        
        //zero-in speed to restrict camera movement without mouse button clicked
        freeLook.m_XAxis.m_MaxSpeed = 0f;
        freeLook.m_YAxis.m_MaxSpeed = 0f;
	}
	
	
	void Update () {
        //ScrollWheel delta affects FOV: this is used for ZOOM imitation(caution: dirty workaround). 
        //Default threshold: 15<FOV<40
        if (Input.GetAxis("Mouse ScrollWheel") != 0f) // forward
        {
            if (freeLook.m_Lens.FieldOfView > 15f && freeLook.m_Lens.FieldOfView < 40f)
           freeLook.m_Lens.FieldOfView  += Input.GetAxis("Mouse ScrollWheel");
            freeLook.m_Lens.FieldOfView = Mathf.Clamp(freeLook.m_Lens.FieldOfView, 15.1f, 39.9f);
        }

        //Next code allows for camera orbiting ONLY on RMB_down by affecting movespeeds
        if (Input.GetMouseButton(1) == false)
        {
            freeLook.m_XAxis.m_MaxSpeed = 0f;
            freeLook.m_YAxis.m_MaxSpeed = 0f;
        }
        else
        {
            freeLook.m_XAxis.m_MaxSpeed = 250f;
            freeLook.m_YAxis.m_MaxSpeed = 2f;
        }
        
    }
}
