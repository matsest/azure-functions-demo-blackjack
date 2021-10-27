#func azure functionapp publish $NAME

Publish-AzWebApp -ResourceGroupName $RgName -Name $FunctionAppName -ArchivePath $ArchivePath