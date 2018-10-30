SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[parameters]
(
  [id] [int] IDENTITY(1,1) NOT NULL,
  [name] [nvarchar](255) NULL,
  [display_name] [nvarchar](255) NULL,
  [type] [nvarchar](255) NULL,
  [length] [int] NULL,
  [function_id] [int] NULL,
  [created_at] [datetimeoffset](7) NULL,
  [updated_at] [datetimeoffset](7) NULL,
  [deleted_at] [datetimeoffset](7) NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[parameters] ADD PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
