PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE EquipmentInfo ("Name","Description","InventoryType","EquipSlot","Quality","Equipable_Left","Equipable_Right","AddStats");
INSERT INTO EquipmentInfo VALUES ("Torch","A torch to light up the darkness.","WeaponTorch","LeftHand","Common","TRUE","FALSE",NULL);
INSERT INTO EquipmentInfo VALUES ("Torch","A torch to light up the darkness.","WeaponTorch","LeftHand","Common","TRUE","FALSE",NULL);
INSERT INTO EquipmentInfo VALUES ("Torch","A torch to light up the darkness.","WeaponTorch","LeftHand","Common","TRUE","FALSE",NULL);
INSERT INTO EquipmentInfo VALUES ("Torch","A torch to light up the darkness.","WeaponTorch","LeftHand","Common","TRUE","FALSE",NULL);
COMMIT;
