SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[applications]
(
  [id] [int] IDENTITY(1,1) NOT NULL,
  [name] [nvarchar](255) NULL,
  [host] [nvarchar](255) NULL,
  [port] [int] NULL,
  [username] [nvarchar](255) NULL,
  [password] [nvarchar](255) NULL,
  [db_name] [nvarchar](255) NULL,
  [created_at] [datetimeoffset](7) NULL,
  [updated_at] [datetimeoffset](7) NULL,
  [deleted_at] [datetimeoffset](7) NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[applications] ADD PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
